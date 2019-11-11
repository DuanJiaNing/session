package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"session/service"
)

const (
	//port = ":10021"
	port = "127.0.0.1:10021"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	service.RegisterSessionServiceServer(s)

	// Register reflection service on gRPC server.
	// reflection.Register(s)
	log.Printf("rpc server listeing on %v", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
