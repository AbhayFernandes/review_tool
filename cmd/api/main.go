package main

import (
	"context"
	"log"
	"net"

	"github.com/AbhayFernandes/review_tool/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
    proto.UnimplementedReviewServiceServer
}

func (s *server) SayHello(
    ctx context.Context, input *proto.HelloRequest,
) (*proto.HelloReply, error) {
    return &proto.HelloReply{
        Message: "Hello, " + input.Name,
    }, nil
}

func main() {
    log.Println("Starting API Service")
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatalln("failed to create listener: ", err)
    }

    s := grpc.NewServer()
    // Make sure this is removed for prod!
    reflection.Register(s)

    proto.RegisterReviewServiceServer(s, &server{})
    if err := s.Serve(listener); err != nil {
        log.Fatalln("failed to serve:", err)
    }
}

