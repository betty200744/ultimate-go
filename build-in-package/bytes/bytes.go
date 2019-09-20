package main

import (
	"bytes"
	"fmt"
)

func main() {

	b := []byte(`ab`)
	bb := []byte{'a', 'b'}
	buffer := bytes.NewBuffer(b)
	fmt.Println(b)
	fmt.Println(bb)
	fmt.Println(buffer)
}
