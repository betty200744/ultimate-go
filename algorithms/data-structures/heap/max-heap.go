package heap

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
type MaxHeap struct {
	Slice    []int
	HeapSize int
}

// left child, right child 与parent 比较，找出最大的
func (h *MaxHeap) MaxHeapify(i int) { // i is array index
	l, r := 2*i+1, 2*i+2 // l is left child node, r is right child node
	max := i
	if l < h.HeapSize && h.Slice[l] > h.Slice[max] {
		max = l
	}
	if r < h.HeapSize && h.Slice[r] > h.Slice[max] {
		max = r
	}
	if max != i {
		h.Slice[i], h.Slice[max] = h.Slice[max], h.Slice[i]
		h.MaxHeapify(max)
	}
}
