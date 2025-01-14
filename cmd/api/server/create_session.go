package server

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"
	"time"

	"github.com/AbhayFernandes/review_tool/pkg/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) RequestNonce(ctx context.Context, input *proto.CreateSessionReply) (*proto.CreateSessionReply, error) {
	log.Println("Recieved nonce generation request")
	db := s.Client.Database("review_service")

	nonce := generateNonce()
	for checkIfNonceExists(db, nonce) {
		log.Println("Nonce already exists, generating new nonce")
		nonce = generateNonce()
	}

	_, err := insertNonceIntoDatabase(db, nonce)

	if err != nil {
		return nil, err
	}

	return &proto.CreateSessionReply{}, nil
}

func generateNonce() string {
	nonceBytes := make([]byte, 16)
	_, err := rand.Read(nonceBytes)
	if err != nil {
		log.Println("Failed to generate nonce")
	}

	return base64.URLEncoding.EncodeToString(nonceBytes)
}

func checkIfNonceExists(db *mongo.Database, nonce string) bool {
	collect := db.Collection("nonces")
	res := collect.FindOne(context.Background(), bson.D{
		bson.E{Key: "code", Value: nonce},
	})

	if res.Err() != nil {
		return false
	}

	return true
}

func createSession(db *mongo.Database) (string, error) {
	collect := db.Collection("sessions")

	res, err := collect.InsertOne(context.Background(), bson.D{
		bson.E{Key: "user", Value: PENDING},
		bson.E{Key: "status", Value: PENDING},
		bson.E{Key: "createdAt", Value: primitive.NewDateTimeFromTime(time.Now())},
	})

	if err != nil {
		log.Println("Failed to create session")
		return "", err
	}

	id := res.InsertedID.(bson.ObjectID).String()

	return id, nil
}

func insertNonceIntoDatabase(db *mongo.Database, nonce string) (*mongo.InsertOneResult, error) {
	collect := db.Collection("nonces")
	res, err := collect.InsertOne(context.Background(), bson.D{
		bson.E{Key: "code", Value: nonce},
		bson.E{Key: "status", Value: PENDING},
		bson.E{Key: "createdAt", Value: primitive.NewDateTimeFromTime(time.Now())},
	})

	if err != nil {
		// TODO: Change this to be graceful handling
		return nil, status.Errorf(codes.Internal, "inserting into database collection failed: \n")
	}

	return res, nil
}
