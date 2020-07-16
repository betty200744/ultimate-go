package maxheap

import (
	"fmt"
	"testing"
)

func TestBuildMaxHeap(t *testing.T) {
	arr := []int{7, 2, 6, 9, 3, 5, 1, 8}
	maxHeap := BuildMaxHeap(arr)
	fmt.Println(maxHeap.Items)
}
func TestMaxHeap_Add(t *testing.T) {
	arr := []int{7, 2, 6, 9, 3, 5, 1, 8}
	maxHeap := BuildMaxHeap(arr)
	fmt.Println(maxHeap.Items)
	maxHeap.Add(88)
	fmt.Println(maxHeap.Items)
}
func TestMaxHeap_ExtractMax(t *testing.T) {
	arr := []int{7, 2, 6, 9, 3, 5, 1, 8}
	maxHeap := BuildMaxHeap(arr)
	fmt.Println(maxHeap.Items)
	maxHeap.ExtractMax()
	fmt.Println(maxHeap.Items)
}
func TestHeapSort(t *testing.T) {
	arr := []int{7, 2, 6, 9, 3, 5, 1, 8}
	HeapSort(arr)
}
