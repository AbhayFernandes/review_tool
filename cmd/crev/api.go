package main

import (
	"log"
	"os"
	"time"

	"crypto/tls"

	"github.com/AbhayFernandes/review_tool/pkg/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func getClient(serverAddr *string) (proto.ReviewServiceClient, context.Context, *grpc.ClientConn, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: false})

	var conn *grpc.ClientConn
	if os.Getenv("STAGE") == "devo" {
		conn, _ = grpc.NewClient(*serverAddr, grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		))
	} else {
		conn, _ = grpc.NewClient(*serverAddr, grpc.WithTransportCredentials(creds))
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
