package main

import (
	"net"
)

func main() {
	// ListenUDP, DialUDP
	udpAddr := &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8887,
	}
	net.ListenUDP("udp", udpAddr)
}
