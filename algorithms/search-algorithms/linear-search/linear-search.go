package main

import (
	"fmt"
	"ultimate-go/algorithms/utils"
)

/*
查找数组中某个value, 天天用到
*/
func linearSearch(arr []int, n int) {
	for i, i2 := range arr {
		if i2 == n {
			fmt.Println(i)
		}
	}
}
func main() {
	arr := utils.GetArrayOfLenAndSize(5, 100)
	linearSearch(arr, arr[3])
}
