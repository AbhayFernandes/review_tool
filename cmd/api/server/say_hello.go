package server

import (
	"context"
	"log"

	"github.com/AbhayFernandes/review_tool/pkg/proto"
)

func (s *Server) SayHello(ctx context.Context, input *proto.HelloRequest) (*proto.HelloReply, error) {
    log.Println("Received: ", input.Name);
	return &proto.HelloReply{Message: "Hello, " + input.Name}, nil
}
