package main

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

type server struct {
	proto.UnimplementedReviewServiceServer;
    *mongo.Client
}

const (
    PENDING = "pending"
    APPROVED = "approved"
    MERGED = "merged"
    CANCELLED = "cancelled"
)


func (s *server) SayHello(
	ctx context.Context, input *proto.HelloRequest,
) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "Hello, " + input.Name,
	}, nil
}

type UserPubKey struct {
	User string             `bson:"user"`
	Sig  string             `bson:"sig"`
}

func (s *server) UploadDiff(
    ctx context.Context, input *proto.UploadDiffRequest,
) (*proto.UploadDiffReply, error) {
    log.Println("Recieved an upload diff request from: " + input.User)

    db := s.Client.Database("review_service")
    collect := db.Collection("diffs")
    sigCollection := db.Collection("sigs")

    var result UserPubKey
    err := sigCollection.FindOne(context.Background(), bson.M{"user": input.User}).Decode(&result)

    if err != nil {
        return nil, status.Errorf(codes.Internal, "You did not supply a valid, registered user. " + err.Error())
    }

    res, err := collect.InsertOne(ctx, bson.D{
        bson.E{Key: "user", Value: input.User},
        bson.E{Key: "state", Value: PENDING},
    })

    if err != nil {
        // TODO: Change this to be graceful handling
        return nil, status.Errorf(codes.Internal, "inserting into database collection failed: \n" + err.Error())
    }

    metadata, ok := metadata.FromIncomingContext(ctx)
    if (!ok) {
        return nil, status.Errorf(codes.PermissionDenied, "You did not supply any headers for authentication")
    }

    ssh_header := metadata.Get("ssh_sig")
    if (len(ssh_header) == 0) {
        return nil, status.Errorf(codes.PermissionDenied, "You did not provide an ssh signature.")
    }

    if (!ssh.Verify(ssh_header[0], input.Diff, result.Sig)) {
        return nil, status.Errorf(codes.PermissionDenied, "Your ssh signature was incorrect.")
    }

    id_i := res.InsertedID
    id, _ := id_i.(bson.ObjectID)

    fo, err := os.Create(getDiffFile(id))
    if (err != nil) {
        // TODO: Change this to be graceful handling
        log.Fatalln("Could not create diff file")
    }
    fo.Write([]byte(input.Diff))
    
    fo.Close()

    log.Println("Completed diff upload for: " + input.User + ", id: " + id.String())    

    return &proto.UploadDiffReply{}, nil
}

func getDiffFile(id bson.ObjectID) string {
    return "/diffs/" + id.Hex() + ".diff"
}
