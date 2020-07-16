package minheap

import "gobyexample/algorithms/data-structures/heap"

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
	if len(arr) > 0 {

	}
	return h
}
func (h *MinHeap) MinHeapifyDown(parentIndex int) {

}
func (h *MinHeap) MinHeapifyUp(childIndex int) {

}
func (h *MinHeap) Insert(item int) {

}
func (h *MinHeap) ExtractMin() int {
	return 0
}
func (h *MinHeap) HeapSort(items []int) {

}
