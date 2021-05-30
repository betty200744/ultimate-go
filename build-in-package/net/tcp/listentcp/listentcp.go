package main

import (
	"net"
)

func main() {
	// ListenTCP, DialTCP
	tcpAddr := &net.TCPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8889,
	}
	net.ListenTCP("tcp", tcpAddr)
}
