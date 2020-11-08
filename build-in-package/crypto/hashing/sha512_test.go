package hashing

import (
	"crypto/sha512"
	"fmt"
	"testing"
)

func Test_sha512(t *testing.T) {
	h := sha512.New()
	h.Write([]byte("abc"))
	fmt.Printf("%x \n", h.Sum(nil))
}
