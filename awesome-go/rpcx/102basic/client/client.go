package main

import (
	"context"
	"flag"
	"fmt"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	// discovery
	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	// xclient
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	//args
	args := &example.Args{
		A: 1,
		B: 2,
	}
	reply := &example.Reply{}
	// reply
	// call
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*reply)
}
