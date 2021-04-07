package main

import (
	"flag"
	"fmt"
)

var ip = flag.String("ip", "127.0.0.1", "-ip : address")
var port = flag.String("port", "80", "-port : port")
var username = flag.String("username", "admin", "-username : username")
var password = flag.String("password", "123456", "-password : password")

// go run flag.go -ip='192.168.2.111' -port=3308 -username='betty' -password='1q2w3e'
func main() {
	flag.Parse()
	fmt.Println(flag.Args())
	fmt.Println(*ip)
	fmt.Println(*port)
	fmt.Println(*username)
	fmt.Println(*password)
}
