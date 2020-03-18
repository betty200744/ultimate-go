package main

import "fmt"

func main() {
	// make , empty slices,  twoD slices, set, get, len, cap, append, copy, slice , for ,
	s := make([]int, 0, 10)
	ss := make([]int, 10)
	sss := make([][]int, 0, 10)
	s8 := []int{1, 2, 3, 4, 5}
	s8c := s8[:]
	fmt.Println("s8c", s8c)
	s = append(s, 9)
	s[0] = 1
	copy(ss, s)
	for e := range s {
		fmt.Println(e)
	}

	for key, value := range s {
		fmt.Println(key, value)
	}

	fmt.Println(s, s[0], len(s), cap(s), ss, s[:2], s[1:2], s[1:], sss)
}
