package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/AbhayFernandes/review_tool/pkg/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
    serverAddr := flag.String(
        "server", "localhost:8080",
        "The server address in the form of host:port",
    )
    flag.Parse()

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Update this to use TLS when moving to PROD!
    conn, err := grpc.NewClient(*serverAddr, grpc.WithTransportCredentials(
        insecure.NewCredentials(),
    ))

    if err != nil {
        log.Fatalln("fail to dial: ", err)
    }
    defer conn.Close()

    client := proto.NewReviewServiceClient(conn)

    res, err := client.SayHello(ctx, &proto.HelloRequest{
        Name: "Test!",
    })
    if (err != nil) {
        log.Fatalln("error sending request: ", err)
    }

    fmt.Println("Response: ", res)
}

