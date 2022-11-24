package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	pb "proto"

	"google.golang.org/grpc"
)

const (
	DefaultAddress = ":10001"
)

const (
	greeter_server_address = ":50051"
)

func getConnByAddress(address string) *grpc.ClientConn {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	conn := getConnByAddress(greeter_server_address)
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	return client.SayHello(ctx, req)
}

func greeterHandle(w http.ResponseWriter, r *http.Request) {
	name := "World"
	p := r.URL.Query().Get("name")
	if p != "" {
		name = p
	}
	res, err := SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("greeter: %v", res.Message)
	b, _ := json.Marshal(res)
	w.Write(b)
}

func main() {
	http.HandleFunc("/", greeterHandle)
	http.ListenAndServe(DefaultAddress, nil)
}
