package main

import (
	"context"
	"log"
	"net"

	"github.com/AbhayFernandes/review_tool/pkg/proto"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting API Service")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to create listener: ", err)
	}

	s := grpc.NewServer()

    client, err := mongo.Connect(options.Client().ApplyURI("mongodb://db:27017"))
    if (err != nil) {
        log.Fatalln("Failed to connect to mongo")
    }

    if err := client.Ping(context.TODO(), nil); err != nil {
        log.Fatalln("Could not connect to MongoDB: ", err)
    }


    defer func() {
        if err = client.Disconnect(context.Background()); err != nil {
            log.Fatalln("Failed to disconnect from mongo")
            panic(err)
        }
    }()

	proto.RegisterReviewServiceServer(s, &server{
        proto.UnimplementedReviewServiceServer{},
        client,
    })

	if err := s.Serve(listener); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
