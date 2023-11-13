package minheap

import (
	"fmt"
	"ultimate-go/algorithms/data-structures/heap"
)

/*
* the largest is on the top
* parent is big then child
* root : i = 0
* parent(i):  (i-1)/2
* left(i): 2*i + 1
* right(i): 2*i + 2
* 2^n , n 为level， 2^n为最小size

BuildMinHeap, 创建MinHeap
MinHeapifyUp, 向上对比，传入childIndex, 递归与parent对比，大的做parent
MinHeapifyDown, 向下对比，传入parentIndex, 递归与left and right child对比，大的做parent
Insert(item int), 添加，append到最后，因为新增child所有需要MinHeapifyUp
ExtractMin() int, 删除root， 删除前先与last交换，最后删除lastItem, 因为root变了， 所有需要MinHeapifyDown
HeapSort(), 排序
*/

type MinHeap struct {
	*heap.Heap
}

func BuildMinHeap(arr []int) MinHeap {
	h := MinHeap{&heap.Heap{Items: arr, HeapSize: len(arr)}}
	// len(arr) / 2, 即从倒数第二层开始heapify, 除于2是为了减少heapify，从最下面的node开始往上比较
	for i := len(arr) / 2; i >= 0; i-- {
		// i 与 i的left and right child对比
		h.MinHeapifyDown(i)
	}
	return h
}

// 向下对比，传入parentIndex, 递归与left and right child对比，小的做parent
func (h *MinHeap) MinHeapifyDown(parentIndex int) {
	l, r := h.GetLeftChildIndex(parentIndex), h.GetRightChildIndex(parentIndex)
	min := parentIndex
	if l < h.HeapSize && h.Items[l] < h.Items[min] {
		min = l
	}
	if r < h.HeapSize && h.Items[r] < h.Items[min] {
		min = r
	}
	if min != parentIndex {
		h.Items[parentIndex], h.Items[min] = h.Items[min], h.Items[parentIndex]
		h.MinHeapifyDown(min)
	}
}

// 向上对比，传入childIndex, 递归与parent对比，小的做parent
func (h *MinHeap) MinHeapifyUp(childIndex int) {
	parentIndex := h.GetParentIndex(childIndex)
	if parentIndex >= 0 && h.Items[childIndex] < h.Items[parentIndex] {
		h.Items[childIndex], h.Items[parentIndex] = h.Items[parentIndex], h.Items[childIndex]
		h.MinHeapifyUp(parentIndex)
	}
}
func (h *MinHeap) Insert(item int) {
	h.Items = append(h.Items, item)
	h.HeapSize++
	h.MinHeapifyUp(h.HeapSize - 1)
}

// 删除root， 删除前先与last交换，最后删除lastItem, 因为root变了， 所有需要MaxHeapifyDown
func (h *MinHeap) ExtractMin() int {
	if h.HeapSize == 0 {
		panic("no items")
	}
	// 取出最小的，即最顶端的
	min := h.Items[0]
	// 用最后的child填充h.Items[0]
	h.Items[0] = h.Items[h.HeapSize-1]
	// 删除最后的child
	h.Items = h.Items[:h.HeapSize-1]
	// 删除最后的child，此时HeapSize - 1
	h.HeapSize--
	// 重排index 0
	h.MinHeapifyDown(0)
	return min
}
func HeapSort(items []int) {
	minHeap := BuildMinHeap(items)
	fmt.Println(minHeap.Items)
	for i := len(minHeap.Items) - 1; i > 0; i-- {
		minHeap.Items[0], minHeap.Items[i] = minHeap.Items[i], minHeap.Items[0]
		minHeap.HeapSize--
		minHeap.MinHeapifyDown(0)
	}
	fmt.Println(minHeap.Items)
}
