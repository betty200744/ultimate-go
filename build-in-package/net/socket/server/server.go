package main

import (
	"fmt"
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", "127.0.0.1:8000")
	for {
		conn, _ := ln.Accept()
		rec := make([]byte, 10)
		conn.Read(rec)
		fmt.Println(fmt.Sprintf("get client id: %v, msg %v", conn.RemoteAddr().String(), string(rec[:])))
		conn.Write([]byte(fmt.Sprintf("server recive msg:  %v", string(rec[:]))))
	}
}
