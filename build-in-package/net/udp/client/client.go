package main

import (
	"fmt"
	"net"
)

func main() {
	//  Dial udp
	udpAddrString := ":8887"
	conn, _ := net.Dial("udp", udpAddrString)
	fmt.Println(conn.LocalAddr(), conn.RemoteAddr())
	// read data
	b := make([]byte, 1024)
	conn.Read(b)
	fmt.Println(string(b[:]))
	conn.Close()
}
