package slice

import (
	"fmt"
	"testing"
)

func Test_inspectSlice(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
			s5 := s4[2:]
			fmt.Println(s5)

			// Slice Iteration by range
			for key, value := range s4 {
				fmt.Println("Slice Iteration by range: ", key, value)
			}
			inspectSlice(s4)
		})
	}
}
