package main

import (
	"fmt"
)

func isOverlap(rangeArray [][]int64) bool {
	for _, i2 := range rangeArray {
		for i4 := 1; i4 < len(rangeArray); i4++ {
			startA := i2[0]
			endA := i2[1]
			startB := rangeArray[i4][0]
			endB := rangeArray[i4][1]
			if startA > startB && startA < endB || startB > startA && startB < endA {
				return true
			}
		}
	}
	return false
}

func main() {

	rangeArray1 := [][]int64{{0, 3}, {4, 6}}
	rangeArray2 := [][]int64{{2, 5}, {1, 5}, {6, 9}}
	if isOverlap(rangeArray1) {
		fmt.Println("1 ")
	}
	if isOverlap(rangeArray2) {
		fmt.Println("2")
	}
}
