package main

import (
	"context"
	"log"
	"os"

	"github.com/AbhayFernandes/review_tool/pkg/proto"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
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

func (s *server) UploadDiff(
    ctx context.Context, input *proto.UploadDiffRequest,
) (*proto.UploadDiffReply, error) {
    log.Println("Recieved an upload diff request from: " + input.User)

    db := s.Client.Database("review_service")
    collect := db.Collection("diffs")

    res, err := collect.InsertOne(ctx, bson.D{
        {"user", input.User},
        {"state", PENDING},
    })

    if err != nil {
        // TODO: Change this to be graceful handling
        log.Fatalln("inserting into collection failed: \n" + err.Error())
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
