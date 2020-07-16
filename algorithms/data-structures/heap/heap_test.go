package heap

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {

	m.Run()
}
func TestHeap_GetLeftChildIndex(t *testing.T) {
	heap := new(Heap)
	heap.Items = []int{0, 1, 2, 3, 4}
	heap.HeapSize = len(heap.Items)
	fmt.Println(heap.GetLeftChildIndex(1))
	fmt.Println(heap.GetRightChildIndex(1))
	fmt.Println(heap.GetParentIndex(4))
	fmt.Println(heap.GetParentIndex(9))
	fmt.Println(heap.GetParentIndex(0))
}
