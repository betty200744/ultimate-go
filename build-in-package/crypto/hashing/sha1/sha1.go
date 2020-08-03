package sha1

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func NewSha1() {
	h := sha1.New()
	io.WriteString(h, "abc")
	fmt.Println(string(h.Sum(nil)))
}
