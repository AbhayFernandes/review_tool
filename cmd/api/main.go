package main

import (
	"log"
	"net"
	"os"

	"github.com/AbhayFernandes/review_tool/cmd/api/server"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.Println("Starting API Service")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to create listener: ", err)
	}

	s, close := server.CreateServer()

    if (os.Getenv("STAGE") == "devo") {
        reflection.Register(s)
    }

	defer close()

	if err := s.Serve(listener); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
