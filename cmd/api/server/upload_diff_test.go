package server

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/AbhayFernandes/review_tool/pkg/proto"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"google.golang.org/grpc/metadata"
)

var dbClient *mongo.Client

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// pull mongodb docker image for version 5.0
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "latest",
		Env:        []string{},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	err = pool.Retry(func() error {
		dbClient, err = mongo.Connect(options.Client().ApplyURI(
			fmt.Sprintf("mongodb://localhost:%s", resource.GetPort("27017/tcp")),
		))
		if err != nil {
			return err
		}
		return dbClient.Ping(context.TODO(), readpref.Nearest())
	})

	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	defer func() {
		// disconnect mongodb client
		if err = dbClient.Disconnect(context.Background()); err != nil {
			log.Fatalf("Could not disconnect from db: %s", err.Error())
		}

		// When you're done, kill and remove the container
		if err = pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}()

	db := dbClient.Database("review_service")
	ctx := context.Background()
	_ = db.CreateCollection(ctx, "sigs")
	sigs := db.Collection("sigs")
	sigs.InsertOne(ctx, bson.D{
		bson.E{Key: "user", Value: "ferna355"},
		bson.E{Key: "sig", Value: "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKC0UcZkZgG6kxqK2/QdWu4oyjiUC9yghwW3Rgav1iqH"},
	})

	m.Run()
}

func TestFetchUserPubKey(t *testing.T) {
	res, err := fetchUserPubKey(dbClient.Database("review_service"), "ferna355")
	if err != nil {
		t.Fatalf("Unable to get the user public key: %s", err)
	}

	if res.Sig != "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKC0UcZkZgG6kxqK2/QdWu4oyjiUC9yghwW3Rgav1iqH" {
		t.Fatalf("The key retrieved is not correct")
	}
}

func TestInsertDataToDatabase(t *testing.T) {
	input := proto.UploadDiffRequest{
		User: "ferna355",
		Diff: "test",
	}
	res, err := insertDataToDatabase(dbClient.Database("review_service"), &input)
	if err != nil {
		t.Fatalf("Unable to insert into database: %s", err)
	}

	if !res.Acknowledged {
		t.Fatalf("The write failed.")
	}
}

func TestUploadDiff_EmptyDiff(t *testing.T) {
	server := Server{Client: dbClient}
	ctx := context.Background()
	input := &proto.UploadDiffRequest{
		User: "ferna355",
		Diff: "",
	}

	_, err := server.UploadDiff(ctx, input)
	if err == nil {
		t.Errorf("Expected error for empty diff, got nil")
	}
}

func TestUploadDiff_InvalidUser(t *testing.T) {
	server := Server{Client: dbClient}
	ctx := context.Background()
	input := &proto.UploadDiffRequest{
		User: "invalid_user",
		Diff: "test",
	}

	_, err := server.UploadDiff(ctx, input)
	if err == nil {
		t.Errorf("Expected error for invalid user, got nil")
	}
}

func TestUploadDiff_InvalidSSHSignature(t *testing.T) {
	server := Server{Client: dbClient}
	ctx := context.Background()
	input := &proto.UploadDiffRequest{
		User: "ferna355",
		Diff: "test",
	}

	md := metadata.Pairs("ssh_sig", "invalid_signature")
	ctx = metadata.NewIncomingContext(ctx, md)

	_, err := server.UploadDiff(ctx, input)
	if err == nil {
		t.Errorf("Expected error for invalid SSH signature, got nil")
	}
}

func TestUploadDiff_MissingSSHSignatureHeader(t *testing.T) {
	server := Server{Client: dbClient}
	ctx := context.Background()
	input := &proto.UploadDiffRequest{
		User: "ferna355",
		Diff: "test",
	}

	_, err := server.UploadDiff(ctx, input)
	if err == nil {
		t.Errorf("Expected error for missing SSH signature header, got nil")
	}
}

func TestUploadDiff_MissingSSHSignature(t *testing.T) {
	server := Server{Client: dbClient}
	ctx := context.Background()
	input := &proto.UploadDiffRequest{
		User: "ferna355",
		Diff: "test",
	}

	md := metadata.Pairs("ssh_sig", "")
	ctx = metadata.NewIncomingContext(ctx, md)

	_, err := server.UploadDiff(ctx, input)
	if err == nil {
		t.Errorf("Expected error for missing SSH signature, got nil")
	}
}
