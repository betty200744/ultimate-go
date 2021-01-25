package main

import (
	"github.com/smallnest/rpcx/server"

	"gobyexample/awesome-go/rpcx/arith"
)

func main() {
	s := server.NewServer()
	s.Register(new(arith.Arith), "")
	s.Serve("tcp", "localhost:8972")
}
