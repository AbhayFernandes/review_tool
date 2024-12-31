package main

import (
	"context"
	"testing"

	"github.com/AbhayFernandes/review_tool/pkg/proto"
)

func TestSayHello_ValidName(t *testing.T) {
	client, ctx, conn, cancel := getClient("localhost:8080")
	defer conn.Close()
	defer cancel()

	resp := sayHello(client, ctx)
	expected := "Hello, Test!"

	if resp != expected {
		t.Errorf("sayHello() = %v; want %v", resp, expected)
	}
}

func TestSayHello_EmptyName(t *testing.T) {
	client, ctx, conn, cancel := getClient("localhost:8080")
	defer conn.Close()
	defer cancel()

	res, err := client.SayHello(ctx, &proto.HelloRequest{
		Name: "",
	})

	if err != nil {
		t.Errorf("sayHello() got unexpected error: %v", err)
	}

	expected := "Hello, "
	if res.GetMessage() != expected {
		t.Errorf("sayHello() = %v; want %v", res.GetMessage(), expected)
	}
}

func TestSayHello_LongName(t *testing.T) {
	client, ctx, conn, cancel := getClient("localhost:8080")
	defer conn.Close()
	defer cancel()

	res, err := client.SayHello(ctx, &proto.HelloRequest{
		Name: "ThisIsAVeryLongNameThatExceedsNormalLength",
	})

	if err != nil {
		t.Errorf("sayHello() got unexpected error: %v", err)
	}

	expected := "Hello, ThisIsAVeryLongNameThatExceedsNormalLength"
	if res.GetMessage() != expected {
		t.Errorf("sayHello() = %v; want %v", res.GetMessage(), expected)
	}
}
