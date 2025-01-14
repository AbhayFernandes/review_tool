package server

import (
	"context"
	"log"
	"os"

	"github.com/AbhayFernandes/review_tool/pkg/proto"
	"github.com/AbhayFernandes/review_tool/pkg/ssh"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type UserPubKey struct {
	User string `bson:"user"`
	Sig  string `bson:"sig"`
}

func (s *Server) UploadDiff(ctx context.Context, input *proto.UploadDiffRequest) (*proto.Unit, error) {
	log.Println("Recieved an upload diff request from: " + input.User)
	db := s.Client.Database("review_service")

	result, err := fetchUserPubKey(db, input.User)
	if err != nil {
		return nil, err
	}

	err = verifySsh(ctx, input.Diff, result.Sig)
	if err != nil {
		return nil, err
	}

	res, err := insertDiffToDatabase(db, input)
	if err != nil {
		return nil, err
	}

	err = saveDiff(res, input.Diff)
	if err != nil {
		return nil, err
	}

	return &proto.Unit{}, nil
}

func fetchUserPubKey(db *mongo.Database, user string) (UserPubKey, error) {
	sigCollection := db.Collection("sigs")

	var result UserPubKey
	err := sigCollection.FindOne(context.Background(), bson.M{"user": user}).Decode(&result)

	if err != nil {
		return result, status.Errorf(codes.Internal, "You did not supply a valid, registered user. ")
	}

	return result, nil
}

func insertDiffToDatabase(db *mongo.Database, input *proto.UploadDiffRequest) (*mongo.InsertOneResult, error) {
	collect := db.Collection("diffs")
	res, err := collect.InsertOne(context.Background(), bson.D{
		bson.E{Key: "user", Value: input.User},
		bson.E{Key: "state", Value: PENDING},
	})

	if err != nil {
		// TODO: Change this to be graceful handling
		return nil, status.Errorf(codes.Internal, "inserting into database collection failed: \n")
	}

	return res, nil
}

func verifySsh(ctx context.Context, diff, sig string) error {
	metadata, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.PermissionDenied, "You did not supply any headers for authentication")
	}

	ssh_header := metadata.Get("ssh_sig")
	if len(ssh_header) == 0 {
		return status.Errorf(codes.PermissionDenied, "You did not provide an ssh signature.")
	}

	if !ssh.Verify(ssh_header[0], diff, sig) {
		return status.Errorf(codes.PermissionDenied, "Your ssh signature was incorrect.")
	}

	return nil
}

func saveDiff(res *mongo.InsertOneResult, diff string) error {
	id_i := res.InsertedID
	id, _ := id_i.(bson.ObjectID)

	fo, err := os.Create(getDiffFile(id))
	if err != nil {
		// TODO: Change this to be graceful handling
		return status.Errorf(codes.Internal, "Failed to save your diff file: %s", err)
	}
	fo.Write([]byte(diff))
	fo.Close()

	log.Println("Completed diff upload. id: " + id.String())

	return nil
}

func getDiffFile(id bson.ObjectID) string {
	return "/diffs/" + id.Hex() + ".diff"
}
