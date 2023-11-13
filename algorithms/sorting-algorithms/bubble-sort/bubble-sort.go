package main

import (
	"fmt"

	"ultimate-go/algorithms/utils"
)

/*
 Bubble sort， 也叫sinking sort下沉算法
性能：
Worst case performance O(n^2)，n * n,  n个for循环
Best case performance O(n) ， n, 一个for循环
Average case performance O(n^2)
*/

// 最基础版本
func bubbleSort1(arr []int) {
	// 每一个i循环， 都是为了找到第i为的最小的值， i左边的为已经排序好的， i右边的为还未排序的
	for i := 0; i < len(arr); i++ {
		// 每一个j循环, 从左到右排序，相邻两两对比， 后面大于前面，则两两交换，
		for j := 0; j < len(arr)-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	fmt.Println("最基础版本， bubbleSort1 after: ", arr)
}

// 优化版本一， j循环减少i次循环
func bubbleSort2(arr []int) {
	// 每一个i循环， 都是为了找到第i为的最小的值， i左边的为已经排序好的， i右边的为还未排序的
	for i := 0; i < len(arr); i++ {
		// 每一个j循环, 从左到右排序，相邻两两对比， 后面大于前面，则两两交换，
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	fmt.Println("优化版本一， j循环减少i次循环，bubbleSort2 after: ", arr)
}

// 优化版本二， j循环减少i次循环， i 也减少n次循环
func bubbleSort3(arr []int) {
	// 每一个i循环， 都是为了找到第i为的最小的值， i左边的为已经排序好的， i右边的为还未排序的
	for i := 0; i < len(arr); i++ {
		swag := false
		// 每一个j循环, 从左到右排序，相邻两两对比， 后面大于前面，则两两交换，
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				swag = true
			}
		}
		if swag == false {
			break
		}
	}
	fmt.Println("优化版本二， j循环减少i次循环， i 也减少n次循环, bubbleSort3 after: ", arr)
}

func main() {
	arr := utils.GetArrayOfLenAndSize(8, 10)
	arr1 := append([]int(nil), arr...)
	arr2 := append([]int(nil), arr...)
	arr3 := append([]int(nil), arr...)
	fmt.Println("before: ", arr)
	bubbleSort1(arr1)
	bubbleSort2(arr2)
	bubbleSort3(arr3)
}
