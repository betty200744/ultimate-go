package sha1

import (
	"crypto/sha1"
	"fmt"
)

func NewSha1() {
	h := sha1.New()
	h.Write([]byte("abc"))
	fmt.Printf("%x \n", string(h.Sum(nil)))
}
