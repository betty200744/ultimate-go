package main

import (
	"github.com/smallnest/rpcx/log"
	"github.com/smallnest/rpcx/server"

	"ultimate-go/awesome-go/rpcx/arith"
)

const (
	Address = "localhost:8972"
)

func main() {
	s := server.NewServer()
	s.Register(new(arith.Arith), "")
	log.Infof(Address)
	s.Serve("tcp", Address)
}
