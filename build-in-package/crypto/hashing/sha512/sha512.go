package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	h := sha512.New()
	h.Write([]byte("abc"))
	fmt.Printf("%x \n", h.Sum(nil))
}
