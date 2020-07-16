package heap

/*

* root : i = 0
* parent(i):  (i-1)/2
* left(i): 2*i + 1
* right(i): 2*i + 2
* 2^n , n 为level， 2^n为最小size

 */

type Heap struct {
	Items    []int
	HeapSize int
}

func (h *Heap) GetLeftChildIndex(parentIndex int) int {
	index := 2*parentIndex + 1
	return index
}
func (h *Heap) GetRightChildIndex(parentIndex int) int {
	index := 2*parentIndex + 2
	return index
}
func (h *Heap) GetParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}
func (h *Heap) HasParent(childIndex int) bool {
	return h.GetParentIndex(childIndex) >= 0
}
func (h *Heap) Swap(ia, ib int) {
	h.Items[ia], h.Items[ib] = h.Items[ib], h.Items[ia]
}
