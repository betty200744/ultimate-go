package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	h := sha256.New()
	h.Write([]byte("abc"))
	fmt.Printf("%x \n", h.Sum(nil))
}
