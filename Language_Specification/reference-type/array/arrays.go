package main

import (
	"fmt"
	"sort"
)

func main() {
	//  init array, declare array, initial array, initial two Dimensional array, multi-line initial array
	//  set array, get array, len of array , comparison array, append, sort, iteration array for range
	// Declare and initialize
	var a1 [3]int
	fmt.Println("this is a1: ", a1)
	// Declare twoD array
	var a2 [2][2]int
	fmt.Println("this is a2: ", a2)
	// Declare and initial value, longhand var a1 [n]T = [n]T{V1,V2,...,Vn} , shorthand: a1 =: [n]T{V1, V2}
	a3 := []string{"cc", "dd"}
	fmt.Println("this is a3:", a3)
	// Declare
	var a4 []string //declare array
	fmt.Printf("this is a4: %v , len : %v , cap: %v \n", a4, len(a4), cap(a4))
	a5 := [...]string{"a3", "bbb"}
	fmt.Printf("this is a5: %v , len : %v , cap: %v \n", a5, len(a5), cap(a5))
	a6 := [][]string{{"a1", "a2"}, {"b1", "b2"}}
	fmt.Printf("this is a6: %v , len : %v , cap: %v \n", a6, len(a6), cap(a6))
	// value assignment
	a1[0] = 2
	a2[0][0] = 2
	// Append
	a3 = append(a3, "a1")
	// Sort
	sort.Strings(a3)
	// Copy
	copy(a4, a3)
	// Iterate with range
	for i, a := range a3 {
		fmt.Printf("Iterate with range: index: %d, value: %v \n", i, a)
	}
	// Iterate with traditional style
	for i := 0; i < len(a3); i++ {
		fmt.Printf("Iterate with traditional style: index: %d, value: %v \n", i, a3[i])
	}

	// -----------------------------
	// Contiguous memory allocations
	// -----------------------------
	// Declare an array of 6 strings initialized with values.
	a7 := [6]string{"Annie", "Betty", "Charley", "Doug", "Edward", "Hoanh"}
	fmt.Printf("\n=> Contiguous memory allocations\n")
	for i, v := range a7 {
		fmt.Printf("Value[%s]\tAddress[%p] IndexAddr[%p]\n", v, &v, &a7[i])
	}
}
