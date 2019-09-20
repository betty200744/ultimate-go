package main

import "fmt"

/*
append(s S, x ...T) S  // T is the element type of S

append, append a single value
append, append multi value
append, append a slice
append, return a new one
append, special, []byte, second argument can be string
*/

/*
copy(dst, src []T) int
copy(dst []byte, src string) int
copy, first arg is dst
copy, return copy num
*/

func main() {
	// append
	s0 := []int{0}
	s1 := append(s0, 1)
	s2 := append(s1, 2, 3, 4)
	s3 := append(s2, []int{5, 6}...)
	s4 := append(s3, []int{7, 8}[0:]...)
	fmt.Println(s0)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	// copy
	a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	s := make([]int, 6)
	n1 := copy(s, a[0:])
	fmt.Println(n1, ":", s)
	n2 := copy(s, s[2:])
	fmt.Println(n2, ":", s)
}
