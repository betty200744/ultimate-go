package main

import (
	"context"
	"testing"

	pb "proto"
)

var (
	s *Service
)

func TestMain(m *testing.M) {
	s = new(Service)
	m.Run()
}

func TestService_SayHello(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pb.HelloRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.HelloReply
		wantErr bool
	}{
		{
			name: "SayHello",
			args: args{
				ctx: context.Background(),
				in:  &pb.HelloRequest{Name: "betty"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.SayHello(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("SayHello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("SayHello() got = %v, want %v", got, tt.want)
		})
	}
}
