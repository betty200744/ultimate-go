package main

import (
	"fmt"
	"net"
)

func main() {
	// ListenUDP, DialUDP
	udpAddr := &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8887,
	}
	conn4, _ := net.DialUDP("udp", nil, udpAddr)
	defer conn4.Close()
	fmt.Println(conn4.LocalAddr(), conn4.RemoteAddr())
}
