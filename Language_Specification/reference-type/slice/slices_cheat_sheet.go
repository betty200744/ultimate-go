package slice

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
