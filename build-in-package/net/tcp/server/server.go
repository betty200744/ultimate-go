package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		b := make([]byte, 100)
		conn.Read(b)
		fmt.Println(conn.LocalAddr(), conn.RemoteAddr(), string(b[:]))
	}
}
