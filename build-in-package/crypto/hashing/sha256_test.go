package hashing

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func Test_sha256(t *testing.T) {
	h := sha256.New()
	h.Write([]byte("abc"))
	fmt.Printf("%x \n", h.Sum(nil))
}
