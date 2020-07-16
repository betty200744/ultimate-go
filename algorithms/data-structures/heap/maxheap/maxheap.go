package maxheap

import (
	"fmt"
	"gobyexample/algorithms/data-structures/heap"
	"log"
)

/*
* the largest is on the top
* parent is big then child
* root : i = 0
* parent(i):  (i-1)/2
* left(i): 2*i + 1
* right(i): 2*i + 2
* 2^n , n 为level， 2^n为最小size

BuildMaxHeap, 创建MaxHeap
MaxHeapifyUp(childIndex int), 向上对比，传入childIndex, 递归与parent对比，大的做parent
MaxHeapifyDown(parentIndex int), 向下对比，传入parentIndex, 递归与left and right child对比，大的做parent
Insert(item int), 添加，append到最后，因为新增child所有需要MaxHeapifyUp
ExtractMax() int, 删除root， 删除前先与last交换，最后删除lastItem, 因为root变了， 所有需要MaxHeapifyDown
HeapSort(items []int), 排序
*/
type MaxHeap struct {
	*heap.Heap
}

func BuildMaxHeap(arr []int) MaxHeap {
	h := MaxHeap{&heap.Heap{Items: arr, HeapSize: len(arr)}}
	// 整个heap有(len(arr) / 2) 个小heap， 从最下面的node开始往上比较
	for i := len(arr) / 2; i >= 0; i-- {
		// 最每个小heap进行max heapify
		h.MaxHeapifyDown(i)
	}
	return h
}

// 向下对比，传入parentIndex, 递归与left and right child对比，大的做parent
func (h *MaxHeap) MaxHeapifyDown(parentIndex int) { // parentIndex is array index
	l, r := h.GetLeftChildIndex(parentIndex), h.GetRightChildIndex(parentIndex) // l is left child node, r is right child node
	max := parentIndex
	if l < h.HeapSize && h.Items[l] > h.Items[max] {
		max = l
	}
	if r < h.HeapSize && h.Items[r] > h.Items[max] {
		max = r
	}
	if max != parentIndex {
		h.Items[parentIndex], h.Items[max] = h.Items[max], h.Items[parentIndex]
		h.MaxHeapifyDown(max)
	}
}

// 向上对比，传入childIndex, 递归与parent对比，大的做parent
func (h *MaxHeap) MaxHeapifyUp(childIndex int) {
	parent := h.GetParentIndex(childIndex)
	if parent >= 0 && h.Items[childIndex] > h.Items[parent] {
		h.Items[childIndex], h.Items[parent] = h.Items[parent], h.Items[childIndex]
		h.MaxHeapifyUp(parent)
	}
}

// 添加，append到最后，因为新增child所有需要MaxHeapifyUp
func (h *MaxHeap) Insert(item int) {
	h.Items = append(h.Items, item)
	h.HeapSize += 1
	h.MaxHeapifyUp(len(h.Items) - 1)
}

// 删除root， 删除前先与last交换，最后删除lastItem, 因为root变了， 所有需要MaxHeapifyDown
func (h *MaxHeap) ExtractMax() int {
	if len(h.Items) == 0 {
		log.Fatal("no items in the heap")
	}
	maxItem := h.Items[0]
	lastItemIndex := len(h.Items) - 1
	h.Items[0] = h.Items[lastItemIndex]
	h.Items = h.Items[:len(h.Items)-1]
	h.HeapSize -= 1
	h.MaxHeapifyDown(0)
	return maxItem
}
func HeapSort(items []int) {
	h := BuildMaxHeap(items)
	// 从右到左排序， 大的放右边
	for i := len(h.Items) - 1; i > 0; i-- {
		// put max item to last item
		h.Items[0], h.Items[i] = h.Items[i], h.Items[0]
		// rm last item from h, 最大的放最后面，且是sorted partition
		h.HeapSize--
		// 只heapify index 0, 因为整个arr只有index 0变了，不是max heap
		h.MaxHeapifyDown(0)
	}
	fmt.Println(h.Items)
}
