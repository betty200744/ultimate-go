package main

import (
	"fmt"
	"ultimate-go/algorithms/utils"
)

/*
* 插入排序， 如同一手牌， 左起第二张开始与前面的对比， 大则往后挪一位, 最后抽出的牌插入第i位
 */
func insertionSort1(arr []int) {
	// 左起第二张, 从左到右依次选出来比较
	for i := 1; i < len(arr); i++ {
		// 选出的牌
		key := arr[i]
		// "选出的牌"， 依次与"选中的牌的前面的牌"对比，
		j := i - 1
		for ; j >= 0 && arr[j] >= key; j-- {
			// 大则往后挪一位，
			arr[j+1] = arr[j]
		}
		// 最后一个挪动的位置就是key选出的牌应该插入的位置， 因为k--了， 所有要加回来，
		arr[j+1] = key
	}
	fmt.Println(arr)
}

func insertionSort2(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = key
	}
	fmt.Println(arr)
}

func main() {
	arr := utils.GetArrayOfLenAndSize(8, 10)
	insertionSort1(arr)
	insertionSort2(arr)
}
