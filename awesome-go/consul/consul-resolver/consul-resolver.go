package consul_resolver

import (
	"context"
	"log"

	"ultimate-go/awesome-go/consul/register"

	"google.golang.org/grpc/resolver"
)

const schemeName = "consul"

type resolve struct {
	cancelFunc context.CancelFunc
}

// ResolveNow will be skipped due unnecessary in this case
func (r *resolve) ResolveNow(resolver.ResolveNowOptions) {}

// Close closes the resolver.
func (r *resolve) Close() {
	r.cancelFunc()
}

type Builder struct {
}

func (b *Builder) Build(url resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	cli, err := register.NewConsulClient("localhost:8500")
	if err != nil {
		log.Fatalf("failed to get consul client %v", err)
		return nil, err
	}
	items, err := cli.Service(url.Endpoint, "")
	if err != nil {
		return nil, err
	}
	addrs := make([]resolver.Address, 0, len(items))
	for _, addr := range items {
		addrs = append(addrs, resolver.Address{Addr: addr})
	}
	cc.UpdateState(resolver.State{Addresses: addrs})
	return &resolve{}, nil
}
func (b *Builder) Scheme() string {
	return schemeName
}
