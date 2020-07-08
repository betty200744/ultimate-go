package main

import "fmt"

func smallest(arr []int, i int) int {
	min := i
	for j := i + 1; j < len(arr); j++ {
		if arr[j] < arr[min] {
			min = j
		}
	}
	return min
}

func main() {
	arr := []int{12, 34, 54, 2, 3, 1, 11, 8, 7}
	for i := 0; i < len(arr); i++ {
		min := smallest(arr, i)
		arr[min], arr[i] = arr[i], arr[min]
	}
	fmt.Println(arr)
}
