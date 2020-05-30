package main

import (
	"fmt"
)

/*
 Bubble sort， 也叫sinking sort下沉算法
性能：
Worst case performance O(n^2)，n * n,  n个for循环
Best case performance O(n) ， n, 一个for循环
Average case performance O(n^2)
*/
// 从右到左排序，小放前面
func bubbleSort1(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	fmt.Println("bubbleSort1 after: ", arr)
}

// 从左到右排序，小的放前面
func bubbleSort2(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ { // 5个数则比四次
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("bubbleSort2 after: ", arr)
}

// 从左到右排序，小的放前面, 优化版本，即一个for循环里面无交互，则外层for循环可中断
func bubbleSort3(arr []int) {
	for i := 0; i < len(arr); i++ {
		swap := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swap = true
			}
		}
		if swap == false {
			break
		}
	}
	fmt.Println("bubbleSort3 after: ", arr)
}

// 从左到右排序，小的放前面, 次优化版本，即一个for循环里面无交换，则外层for循环可中断，但内循环多运行了n次
func bubbleSort4(arr []int) {
	swap := true
	for swap {
		swap = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swap = true
			}
		}
	}
	fmt.Println("bubbleSort4 after: ", arr)
}

func main() {
	arr := []int{2, 6, 0, 9, 1, 1, 0, 9}
	arr1 := append([]int(nil), arr...)
	arr2 := append([]int(nil), arr...)
	arr3 := append([]int(nil), arr...)
	arr4 := append([]int(nil), arr...)
	fmt.Println("before: ", arr)
	bubbleSort1(arr1)
	bubbleSort2(arr2)
	bubbleSort3(arr3)
	bubbleSort4(arr4)

}
