package main

import (
	"encoding/json"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	a := `client send aaaa`
	b, _ := json.Marshal(a)
	for i := 0; i < 5; i++ {
		conn.Write(b)
	}
}
