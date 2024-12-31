package main

import (
	"context"
	"testing"

	"github.com/AbhayFernandes/review_tool/pkg/proto"
)

func TestSayHello(t *testing.T) {
    server := server{}
    ctx := context.Background()
    input := &proto.HelloRequest{
        Name: "Test",
    }

    resp, err := server.SayHello(ctx, input)
    if err != nil {
        t.Errorf("HelloTest(%v) got unexpected error", err)
    }

    if resp.Message != "Hello, Test" {
        t.Errorf("HelloTest(%v) got unexpected error", err)
    }
}
