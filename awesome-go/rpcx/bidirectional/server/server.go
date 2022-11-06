package main

import (
	"context"
	"fmt"
	"net"
	_ "net/http/pprof"
	"time"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/server"
)

var clientConn net.Conn
var connected = false

type Arith int

func (t *Arith) Mul(ctx context.Context, args *example.Args, reply *example.Reply) error {
	clientConn = ctx.Value(server.RemoteConnContextKey).(net.Conn)
	reply.C = args.A * args.B
	connected = true
	return nil
}

func main() {
	s := server.NewServer()
	s.Register(new(Arith), "")
	go s.Serve("tcp", "localhost:8972")

	for !connected {
		time.Sleep(time.Second)
	}
	for {
		if clientConn != nil {
			err := s.SendMessage(clientConn, "test_service_path", "test_service_method", nil, []byte("abcde"))
			if err != nil {
				fmt.Printf("failed to send messsage to %s: %v\n", clientConn.RemoteAddr().String(), err)
				clientConn = nil
			} else {
				fmt.Printf("send messsage to %s\n", clientConn.RemoteAddr().String())

			}
		}
		time.Sleep(time.Second)
	}
}
