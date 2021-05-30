package main

import (
	"fmt"
	"net"
)

func main() {
	// listen tcp
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		// server read
		b := make([]byte, 1024)
		conn.Read(b)
		fmt.Println(string(b[:]))
		msg := fmt.Sprintf(`s -> c : hello %v`, conn.RemoteAddr())
		fmt.Println(msg)
		conn.Write([]byte(msg))
	}
}
