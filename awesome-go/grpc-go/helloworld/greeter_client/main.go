package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
	pb "ultimate-go/awesome-go/grpc-go/helloworld/helloworld"
)

const (
	address = "localhost:50051"
)

func clientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	return invoker(ctx, method, req, reply, cc, opts...)
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithUnaryInterceptor(clientInterceptor))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := pb.NewGreeterClient(conn)

	for i := 0; i < 10000; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
		defer cancel()
		ctxx := metadata.AppendToOutgoingContext(ctx, "go_client_md_1", "c1", "go_client_md_2", "c2")
		r, err := c.SayHello(ctxx, &pb.HelloRequest{Name: "world"})
		if err != nil {
			log.Printf("could not greet: %v", err)
		} else {
			log.Printf("Greeting: %v, %s", i, r.Message)
		}
		time.Sleep(time.Second)
	}
	/*	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 3)
		defer cancel()
		ctxx := metadata.AppendToOutgoingContext(ctx, "go_client_md_1", "c1", "go_client_md_2", "c2")
		r, err := c.SayHello(ctxx, &pb.HelloRequest{Name: "world"})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.Message)*/
	//conn.Close()
	time.Sleep(time.Second * 20)
}
