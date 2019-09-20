package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	h := md5.New()
	io.WriteString(h, "hello test")
	fmt.Printf("%x", h.Sum(nil))
}
