package main

import (
	"example.com/m/v2/grpc/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// Create gRPC Server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Println("gRPC server is running.")
	pb.RegisterGreeterServer(s, &HelloGrpcService{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
