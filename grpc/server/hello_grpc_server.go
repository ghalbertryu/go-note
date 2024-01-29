package main

import (
	"context"
	"example.com/m/v2/grpc/pb"
)

type HelloGrpcService struct {
	pb.UnimplementedGreeterServer
}

func (m HelloGrpcService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Server say hello to " + in.GetName()}, nil
}
