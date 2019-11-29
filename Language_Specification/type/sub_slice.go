package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	ss := s[1:]
	ss = append(ss, 4)

	for _, v := range ss {
		v += 10 //only change v , not ss
	}

	for i := range ss {
		ss[i] += 10
	}

	fmt.Println(s)
	fmt.Println(ss)
}
