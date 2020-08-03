package md5

import (
	"crypto/md5"
	"fmt"
	"io"
)

func NewMd5() {
	h := md5.New()
	io.WriteString(h, "abc")
	fmt.Printf("%x \n", h.Sum(nil))
}
