package main

import "fmt"

/*
wiki: https://en.wikipedia.org/wiki/Shellsort

*/
func shellSort(arr []int) {
	// gap size，初始一般len(ll) / 2，
	// gap size一般递减1， 也可 / 2, 递减或者除于，最后一个gap需要为1
	for gap := len(arr) / 2; gap > 0; gap /= 2 {
		// i for循环， i 后面的每个都要对比
		for i := gap; i < len(arr); i++ {
			// j for循环， ll[j] 与ll[j-gap]对比，注意j>=gap, 否则out of index
			// j for循环,  注意i > gap很多时， 需要j -= gap, 保证对比不重复
			for j := i; j >= gap && arr[j-gap] > arr[j]; j -= gap {
				arr[j], arr[j-gap] = arr[j-gap], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

func main() {
	arr := []int{12, 34, 54, 2, 3, 1, 11, 8, 7}
	shellSort(arr)
}
