package main

import "fmt"

func main() {
	// string to []byte
	s := []byte("abc")

	// []byte to string
	b := []byte{23, 34}

	fmt.Println(string(b))
	fmt.Println(s)

	// copy
	cd := make([]int, 2)
	copy(cd, []int{3, 4})
	fmt.Println(cd, len(cd))

}
