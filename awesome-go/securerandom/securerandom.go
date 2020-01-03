package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func Base64InBytes(max int) (string, error) {
	b, err := Bytes(maximumBytes(max))
	return base64.StdEncoding.EncodeToString(b), err
}
func maximumBytes(size int) int {
	return int((float64(size) / 4) * 3)
}
func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)

	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

func main() {

	// assume your random string can only be 32 bytes long
	rStr, _ := Base64InBytes(32)

	fmt.Println("this is base64 string: ", rStr) // would print Base64 string with a length of 32

}
