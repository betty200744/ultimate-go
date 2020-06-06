package main

import (
	"fmt"
	"gobyexample/algorithms/utils"
)

/*
* 插入排序， 如同一手牌， 左起第二张开始与前面的对比， 大则往后挪一位, 最后抽出的牌插入第i位
 */
func insertionSort1(arr []int) {
	// 左起第二张, 从左到右依次选出来比较
	for j := 1; j < len(arr); j++ {
		// 选出的牌
		key := arr[j]
		// "选出的牌"， 依次与"选中的牌的前面的牌"对比，
		i := j - 1
		for ; i >= 0 && arr[i] >= key; i-- {
			// 大则往后挪一位，
			arr[i+1] = arr[i]
		}
		// 最后一个挪动的位置就是key选出的牌应该插入的位置， 因为i--了， 所有要加回来，
		arr[i+1] = key
	}
	fmt.Println(arr)
}

func main() {
	arr := utils.GetArrayOfLenAndSize(8, 10)
	insertionSort1(arr)
}
