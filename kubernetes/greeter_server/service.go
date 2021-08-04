package main

import (
	"context"
	"log"

	pb "proto"
)

type Service struct {
	pb.UnimplementedGreeterServer
}

func (s *Service) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received in: %v", in)
	reply := &pb.HelloReply{Message: "hello " + in.Name}
	return reply, nil
}
