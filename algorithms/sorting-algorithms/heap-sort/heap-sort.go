package main

import (
	"fmt"
	"gobyexample/algorithms/data-structures/heap"
	"gobyexample/algorithms/utils"
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
func BuildMaxHeap(arr []int) heap.MaxHeap {
	h := heap.MaxHeap{Slice: arr, HeapSize: len(arr)}
	// 整个heap有(len(arr) / 2) 个小heap， 从最下面的node开始往上比较
	for i := len(arr) / 2; i >= 0; i-- {
		// 最每个小heap进行max heapify
		h.MaxHeapify(i)
	}
	return h
}

func heapSort1(arr []int) {
	h := BuildMaxHeap(arr)
	// 从右到左排序， 大的放右边
	for i := len(h.Slice) - 1; i > 0; i-- {
		// put max item to last item
		h.Slice[0], h.Slice[i] = h.Slice[i], h.Slice[0]
		// rm last item from h, 最大的放最后面，且是sorted partition
		h.HeapSize--
		// 只heapify index 0, 因为整个arr只有index 0变了，不是max heap
		h.MaxHeapify(0)
	}
	fmt.Println(h.Slice)
}

func main() {
	arr := utils.GetArrayOfLenAndSize(8, 10)
	heapSort1(arr)
}
