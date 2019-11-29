package main

import (
	"fmt"
	"math"
	"sort"
)

// func literals
func fooo(a, b int) int {
	return a + b
}

// primary expressions

// selectors expressions

// method expressions

func main() {
	// alternative
	a := 1 | 0
	fmt.Println(a)
	arr := []int{1, 2, 4}
	arr = append(arr, 3)
	for e := range arr {
		fmt.Println(e)
	}
	for k, v := range arr {
		fmt.Println(k, v)
	}
	sort.Ints(arr)
	fmt.Println(arr)
	arrstring := []string{"c", "b", "a"}
	sort.Strings(arrstring)
	fmt.Println(arrstring)
	// Package Qualified identifiers
	fmt.Println(math.Max(2, 4))
	// Composite literals
	fmt.Println([]int{1, 3})

}
