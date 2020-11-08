package utils

import (
	"context"
	"fmt"
	"testing"

	"google.golang.org/grpc/metadata"
)

func TestCopyContext(t *testing.T) {
	ctx := context.Background()
	res := CopyContext(ctx)
	fmt.Println(res)
}

func TestAppendContext(t *testing.T) {
	res := AppendContext(context.Background(), "grpcgateway-user-agen", ",13,,,,,")
	res2 := AppendContext(res, "a", "a")
	fmt.Println(res)
	fmt.Println(res2)
}

func TestAppendToOutgoingContext(t *testing.T) {
	res := metadata.AppendToOutgoingContext(context.Background(), "grpcgateway-user-agen", ",13,,,,,")
	fmt.Println(res)
}
