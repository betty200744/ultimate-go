package main

import (
	"context"
	"fmt"
	"net"

	"gobyexample/utils"
)

func main() {
	var d net.Dialer
	conn, err := d.DialContext(context.Background(), "tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.Write([]byte(utils.RandomString(8)))
	res := make([]byte, 0)
	fmt.Println(string(res[:]))
}
