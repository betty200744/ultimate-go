package main

import "fmt"

// ----------------------
// Declare and initialize
// ----------------------
//  -----
// |  *  | --> | nil | nil | nil | nil | nil |
//  -----      |  0  |  0  |  0  |  0  |  0  |
// |  5  |
//  -----
// |  5  |
//  -----
func inspectSlice(slice []int) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i := range slice {
		fmt.Printf("[%d] %p %v\n", i, &slice[i], slice[i])
	}
}
func main() {
	// make , empty slices,  twoD slices, set, get, len, cap, append, copy, slice , for ,
	// create slice with 10 element, zero value 0
	s1 := make([]int, 10)
	fmt.Printf("s1: %v, len: %v, cap: %v \n", s1, len(s1), cap(s1))
	// create slice with 8 element, and capacity of 10
	s2 := make([]int, 8, 10)
	fmt.Printf("s2: %v, len: %v, cap: %v \n", s2, len(s2), cap(s2))

	// create twoD slices
	s3 := make([][]int, 5, 10)
	fmt.Printf("s3: %v, len: %v, cap: %v \n", s3, len(s3), cap(s3))

	// create and init slice with value
	s4 := []int{1, 2, 3, 4, 5}
	// Slice of slice
	fmt.Println("Slice of slice: ", s4[:])

	// Slice Iteration by range
	for key, value := range s4 {
		fmt.Println("Slice Iteration by range: ", key, value)
	}
	inspectSlice(s4)
}
