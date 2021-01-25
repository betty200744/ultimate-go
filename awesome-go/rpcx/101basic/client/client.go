package main

import (
	"context"
	"fmt"

	"github.com/smallnest/rpcx/client"

	"gobyexample/awesome-go/rpcx/arith"
)

func main() {
	d, _ := client.NewPeer2PeerDiscovery("tcp@localhost:8972", "")
	arithClient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	reply := &arith.Reply{}
	err := arithClient.Call(context.Background(), "Mul", &arith.Args{A: 1, B: 2}, reply)
	if err != nil {
		panic(err)
	}
	fmt.Println("##############", reply)
}
