package main

import (
	"ultimate-go/algorithms/data-structures/heap/maxheap"
)

/*
* Build Heap
* Change to Max Heap
* 递归方式，自己调自己
* complate tree
* Min Heap or Max Heap
* 2^n , n 为level， 2^n为最小size
* root element : arr[0]
* parent node: arr[(i-1)/2]
* left child node: arr[(2*i) + 1]
* right child node: arr[(2*i) + 2]
 */
// 输入arr， 创建Max Heap

func main() {
	arr := []int{7, 2, 6, 9, 3, 5, 1, 8}
	maxheap.HeapSort(arr)
}
