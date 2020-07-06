package main

import "fmt"

/*
wiki: https://en.wikipedia.org/wiki/Shellsort

*/

func shellSort(ll []int) {
	// gap 一般len(ll) / 2
	// gap size一般递减1， 或者 /= 2
	for gap := len(ll) / 2; gap > 0; gap -= 1 {
		for i := gap; i < len(ll); i++ {
			for j := i; j >= gap && ll[j-gap] > ll[j]; j -= gap {
				ll[j], ll[j-gap] = ll[j-gap], ll[j]
			}
		}
	}
	fmt.Println(ll)
}

func main() {
	arr := []int{12, 34, 54, 2, 3, 1, 11, 8, 7}
	shellSort(arr)
}
