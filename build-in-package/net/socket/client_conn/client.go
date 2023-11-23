package main

import (
	"fmt"
	"net"

	"ultimate-go/utils"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.Write([]byte(utils.RandomString(8)))
	res := make([]byte, 1024)
	conn.Read(res)
	fmt.Println(string(res[:]))
}
