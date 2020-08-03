package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	b, _ := aes.NewCipher([]byte("qwertyuiopasdfgd"))
	in := []byte("fdaserearewrearf")
	out := make([]byte, len(in))
	b.Encrypt(out, in)
	fmt.Printf("%x \n", out)
	desD := make([]byte, len(in))
	b.Decrypt(desD, out)
	fmt.Printf("%s \n", desD)
}
