package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen udp
	udpAddrString := &net.UDPAddr{
		IP:   nil,
		Port: 8887,
	}
	l, err := net.ListenUDP("udp", udpAddrString)
	fmt.Println(err)
	// read data
	b := make([]byte, 1024)
	l.Read(b)
	fmt.Println(string(b[:]))
	// write data
	l.Write([]byte(`s -> c: hello`))
	// close
	l.Close()
}
