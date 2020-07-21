package minheap

import (
	"gobyexample/algorithms/data-structures/heap"
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
	h := MinHeap{&heap.Heap{
		Items:    arr,
		HeapSize: len(arr),
	}}
	for i := len(arr) / 2; i >= 0; i-- {
		// 最每个小heap进行max heapify
		h.MinHeapifyDown(i)
	}
	return h
}
func (h *MinHeap) MinHeapifyDown(parentIndex int) {
	l, r := h.GetLeftChildIndex(parentIndex), h.GetRightChildIndex(parentIndex)
	min := parentIndex
	if l < len(h.Items) && h.Items[l] < h.Items[min] {
		min = l
	}
	if r < len(h.Items) && h.Items[r] < h.Items[min] {
		min = r
	}
	if min != parentIndex {
		h.Items[parentIndex], h.Items[min] = h.Items[min], h.Items[parentIndex]
		h.MinHeapifyDown(min)
	}
}
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
	h.MinHeapifyUp(len(h.Items) - 1)
}
func (h *MinHeap) ExtractMin() int {
	if len(h.Items) == 0 {
		panic("no items")
	}
	min := h.Items[0]
	h.Items[0] = h.Items[len(h.Items)-1]
	h.HeapSize--
	h.Items = h.Items[:len(h.Items)-1]
	h.MinHeapifyDown(0)
	return min
}
func HeapSort(items []int) {

}
