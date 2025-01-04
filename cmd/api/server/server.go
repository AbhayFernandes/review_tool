package server

import (
	"github.com/AbhayFernandes/review_tool/pkg/proto"
	"go.mongodb.org/mongo-driver/v2/mongo"

	mong "github.com/AbhayFernandes/review_tool/cmd/api/mongo"
	"google.golang.org/grpc"
)

const (
	PENDING   = "pending"
	APPROVED  = "approved"
	MERGED    = "merged"
	CANCELLED = "cancelled"
)

type Server struct {
	proto.UnimplementedReviewServiceServer
	*mongo.Client
}

func CreateServer() (*grpc.Server, func()) {
	s := grpc.NewServer()
	client, close := mong.GetMongoClient()

	proto.RegisterReviewServiceServer(s, &Server{
		proto.UnimplementedReviewServiceServer{},
		client,
	})

	return s, close
}
