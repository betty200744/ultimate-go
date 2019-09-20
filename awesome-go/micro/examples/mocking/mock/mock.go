package mock

import (
	"context"

	proto "examples/helloworld/proto"
	"github.com/micro/go-micro/client"
)

type mockGreeterService struct {
}

func (m *mockGreeterService) Hello(ctx context.Context, req *proto.HelloRequest, opts ...client.CallOption) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{
		Greeting: "Hello " + req.Name,
	}, nil
}

func NewGreeterService() proto.GreeterService {
	return new(mockGreeterService)
}
