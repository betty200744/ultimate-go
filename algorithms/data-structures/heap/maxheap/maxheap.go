package maxheap

import (
	"fmt"
	"log"
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
	// len(arr) / 2, 即从倒数第二层开始heapify, 除于2是为了减少heapify，从最下面的node开始往上比较
	for i := len(arr) / 2; i >= 0; i-- {
		// i 与 i的left and right child对比
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
	parentIndex := h.GetParentIndex(childIndex)
	if parentIndex >= 0 && h.Items[childIndex] > h.Items[parentIndex] {
		h.Items[childIndex], h.Items[parentIndex] = h.Items[parentIndex], h.Items[childIndex]
		h.MaxHeapifyUp(parentIndex)
	}
}

// 添加，append到最后，因为新增child所有需要MaxHeapifyUp
func (h *MaxHeap) Insert(item int) {
	h.Items = append(h.Items, item)
	h.HeapSize++
	h.MaxHeapifyUp(h.HeapSize - 1)
}

// 删除root， 删除前先与last交换，最后删除lastItem, 因为root变了， 所有需要MaxHeapifyDown
func (h *MaxHeap) ExtractMax() int {
	if h.HeapSize == 0 {
		log.Fatal("no items in the heap")
	}
	// 取出最大的，即最顶端的
	maxItem := h.Items[0]
	// 用最后的child填充h.Items[0]
	h.Items[0] = h.Items[h.HeapSize-1]
	// 删除最后的child
	h.Items = h.Items[:h.HeapSize-1]
	// 删除最后的child，此时HeapSize - 1
	h.HeapSize--
	// 重排index 0
	h.MaxHeapifyDown(0)
	return maxItem
}
func HeapSort(items []int) {
	h := BuildMaxHeap(items)
	// BuildMaxHeap后h.Items[0]最大， 最大放最后面
	// 从后到前即从右到左排序， 大的放右边
	for i := h.HeapSize - 1; i > 0; i-- {
		// put max item to last item
		h.Items[0], h.Items[i] = h.Items[i], h.Items[0]
		// rm last item from h, 最大的放最后面，且是sorted partition
		h.HeapSize--
		// 只heapify index 0, 因为整个arr只有index 0变了，不是max heap
		h.MaxHeapifyDown(0)
	}
	fmt.Println(h.Items)
}
