package main

import (
	"net"
)

func main() {
	tcpAddr := &net.TCPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8889,
	}
	conn3, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer conn3.Close()
}
