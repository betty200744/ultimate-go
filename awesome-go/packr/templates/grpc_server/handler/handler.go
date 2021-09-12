package handler

import (
	"context"

	"ultimate-go/grpc_server/service"
	pb "ultimate-go/proto/helloworld"
)

// Server is used to implement helloworld.GreeterServer.
type GreeterServer struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	ss := service.NewServer()
	return ss.SayHello(ctx, in)
}
