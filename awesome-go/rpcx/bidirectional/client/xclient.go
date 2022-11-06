package main

import (
	"context"
	"fmt"
	"log"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

// xclient NewBidirectionalXClient
func main() {
	ch := make(chan *protocol.Message)
	d, _ := client.NewPeer2PeerDiscovery("tcp@localhost:8972", "")
	xClient := client.NewBidirectionalXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption, ch)
	defer xClient.Close()

	args := &example.Args{A: 10, B: 20}
	reply := &example.Reply{}
	err := xClient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	log.Printf("%d * %d = %d", args.A, args.B, reply.C)

	for msg := range ch {
		fmt.Printf("receive msg from server: %s\n", msg.Payload)
	}
}
