package server

import (
	"context"
	"github.com/AbhayFernandes/review_tool/pkg/proto"
)


func (s *Server) SayHello(ctx context.Context, input *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "Hello, " + input.Name}, nil
}
