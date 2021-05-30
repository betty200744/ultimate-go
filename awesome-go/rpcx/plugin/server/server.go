package main

import (
	"context"
	"log"
	"net"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/server"
)

type ConnectionPlugin struct {
}

func (p *ConnectionPlugin) HandleConnAccept(conn net.Conn) (net.Conn, bool) {
	log.Printf("client %v connected", conn.RemoteAddr().String())
	return conn, true
}

func (p *ConnectionPlugin) HandleConnClose(conn net.Conn) bool {
	log.Printf("client %v closed", conn.RemoteAddr().String())
	return true
}

type Arith struct {
}

func (t *Arith) Mul(ctx context.Context, args example.Args, reply *example.Reply) error {
	reply.C = args.A + args.B
	return nil
}

func main() {
	s := server.NewServer()
	// server plugin
	s.Plugins.Add(&ConnectionPlugin{})
	s.RegisterName("Arith", new(Arith), "")
	err := s.Serve("tcp", "localhost:8972")
	if err != nil {
		panic(err)
	}
}
