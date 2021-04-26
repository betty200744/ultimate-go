package main

import (
	"context"
	"fmt"
	"log"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

// client use RegisterServerMessageChan
func main() {
	c := client.NewClient(client.DefaultOption)
	err := c.Connect("tcp", "localhost:8972")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	args := &example.Args{A: 10, B: 20}
	reply := &example.Reply{}
	err = c.Call(context.Background(), "Arith", "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	log.Printf("%d * %d = %d", args.A, args.B, reply.C)

	ch := make(chan *protocol.Message)
	c.RegisterServerMessageChan(ch)
	for msg := range ch {
		fmt.Printf("receive msg from server: %s\n", msg.Payload)
	}
}
