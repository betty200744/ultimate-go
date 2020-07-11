package main

import (
	"fmt"
)

// 选择一个pivot，

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	lowPart := make([]int, 0)
	highPart := make([]int, 0)
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			lowPart = append(lowPart, arr[i])
		} else {
			highPart = append(highPart, arr[i])
		}
	}
	sorted := QuickSort(lowPart)
	sorted = append(sorted, pivot)
	sorted = append(sorted, QuickSort(highPart)...)
	return sorted
}
func main() {
	arr := []int{7, 2, 6, 9, 3, 5, 1, 8}
	after := QuickSort(arr)
	fmt.Println(after)
}
