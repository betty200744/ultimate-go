package main

import (
	"context"
	"log"
	"net"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/client"
)

var (
	xclient client.XClient
)

type ConnectionPlugin struct {
}

func (p *ConnectionPlugin) ClientConnected(conn net.Conn) (net.Conn, error) {
	log.Printf("server %v connected", conn.RemoteAddr().String())
	return conn, nil
}

func (p *ConnectionPlugin) ClientConnectionClose(conn net.Conn) error {
	log.Printf("server %v closed", conn.RemoteAddr().String())
	return nil
}

func main() {
	d, _ := client.NewPeer2PeerDiscovery("tcp@localhost:8972", "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()
	// client plugins
	plugins := client.NewPluginContainer()
	plugins.Add(&ConnectionPlugin{})
	xclient.SetPlugins(plugins)

	// client call
	args := example.Args{A: 10, B: 20}
	reply := &example.Reply{}
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
