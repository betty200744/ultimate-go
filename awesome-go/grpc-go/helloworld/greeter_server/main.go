//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"ultimate-go/awesome-go/consul/register"
	pb "ultimate-go/awesome-go/grpc-go/helloworld/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	port = ":50053"
)

func serverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	return handler(ctx, req)
}

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Println("##################metadata: ", md)
	log.Printf("Received: %v", in.Name)
	header := metadata.New(map[string]string{"js_server_md_1": "s1", "js_server_md_2": "s2"})
	grpc.SendHeader(ctx, header)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(serverInterceptor))
	pb.RegisterGreeterServer(s, &server{})
	cli, err := register.NewConsulClient("localhost:8500")
	if err != nil {
		log.Fatalf("failed to get consul client %v", err)
	}
	cli.Register("greeter", 50053)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
