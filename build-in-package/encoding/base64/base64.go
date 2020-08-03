package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	msg := "hello, world"
	str := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println(str)
	desMsg, _ := base64.StdEncoding.DecodeString(str)
	fmt.Println(string(desMsg))
}
