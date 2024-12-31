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

func TestSayHello_EmptyName(t *testing.T) {
	server := server{}
	ctx := context.Background()
	input := &proto.HelloRequest{
		Name: "",
	}

	resp, err := server.SayHello(ctx, input)
	if err != nil {
		t.Errorf("HelloTest_EmptyName(%v) got unexpected error", err)
	}

	if resp.Message != "Hello, " {
		t.Errorf("HelloTest_EmptyName(%v) got unexpected error", err)
	}
}

func TestSayHello_LongName(t *testing.T) {
	server := server{}
	ctx := context.Background()
	input := &proto.HelloRequest{
		Name: "ThisIsAVeryLongNameThatExceedsNormalLength",
	}

	resp, err := server.SayHello(ctx, input)
	if err != nil {
		t.Errorf("HelloTest_LongName(%v) got unexpected error", err)
	}

	if resp.Message != "Hello, ThisIsAVeryLongNameThatExceedsNormalLength" {
		t.Errorf("HelloTest_LongName(%v) got unexpected error", err)
	}
}

func TestSayHello_SpecialCharacters(t *testing.T) {
	server := server{}
	ctx := context.Background()
	input := &proto.HelloRequest{
		Name: "!@#$%^&*()_+",
	}

	resp, err := server.SayHello(ctx, input)
	if err != nil {
		t.Errorf("HelloTest_SpecialCharacters(%v) got unexpected error", err)
	}

	if resp.Message != "Hello, !@#$%^&*()_+" {
		t.Errorf("HelloTest_SpecialCharacters(%v) got unexpected error", err)
	}
}
