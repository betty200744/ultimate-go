package utils

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func CopyContext(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.MD{}
	}
	return metadata.NewIncomingContext(ctx, md)
}

func AppendContext(ctx context.Context, key, val string) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.MD{}
	}
	md.Set(key, val)
	return metadata.NewIncomingContext(ctx, md)
}
