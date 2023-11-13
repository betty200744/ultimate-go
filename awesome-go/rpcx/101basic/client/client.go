package main

import (
	"context"
	"fmt"

	"github.com/smallnest/rpcx/client"

	"ultimate-go/awesome-go/rpcx/arith"
)

func main() {
	d, _ := client.NewPeer2PeerDiscovery("tcp@localhost:8972", "")
	xClient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	reply := &arith.Reply{}
	err := xClient.Call(context.Background(), "Mul", &arith.Args{A: 1, B: 2}, reply)
	if err != nil {
		panic(err)
	}
	fmt.Println("##############", reply)
	err = xClient.Call(context.Background(), "Add", &arith.Args{A: 1, B: 2}, reply)
	fmt.Println("##############", reply)
	err = xClient.Call(context.Background(), "Say", "come back home", reply)
	fmt.Println("##############", reply)
}
