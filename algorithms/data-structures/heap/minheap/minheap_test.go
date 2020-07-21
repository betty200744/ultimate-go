package minheap

import (
	"fmt"
	"testing"
)

func TestBuildMinHeap(t *testing.T) {
	arr := []int{7, 2, 6, 9, 3, 5, 1, 8}
	minHeap := BuildMinHeap(arr)
	fmt.Println(minHeap.Items)
	minHeap.Insert(0)
	fmt.Println(minHeap.Items)
	minHeap.ExtractMin()
	fmt.Println(minHeap.Items)
}
func TestMinHeap_HeapSort(t *testing.T) {
	arr := []int{7, 2, 6, 9, 3, 5, 1, 8}
	HeapSort(arr)
}
