package main

import (
	"log"
	"time"

	"github.com/AbhayFernandes/review_tool/pkg/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getClient(serverAddr *string) (proto.ReviewServiceClient, context.Context, *grpc.ClientConn, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	conn, err := grpc.NewClient(*serverAddr, grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))

	if err != nil {
		log.Fatalln("fail to dial: ", err)
	}

	client := proto.NewReviewServiceClient(conn)

	return client, ctx, conn, cancel
}

func sayHello(c proto.ReviewServiceClient, ctx context.Context) string {
	res, err := c.SayHello(ctx, &proto.HelloRequest{
		Name: "Test!",
	})

	if err != nil {
		log.Fatalln("error sending request: ", err)
	}

	return res.GetMessage()
}
