package main

import (
	"context"

	"github.com/AbhayFernandes/review_tool/pkg/proto"
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
