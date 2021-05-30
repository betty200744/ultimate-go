package main

import (
	"fmt"
	"net"
)

func main() {
	// dial tcp
	c, _ := net.Dial("tcp", "localhost:8888")

	// client write
	msg := []byte(`c > s : hello`)
	c.Write(msg)
	// client read
	b := make([]byte, 1024)
	c.Read(b)
	fmt.Println(string(b[:]))
	// close
	c.Close()
}
