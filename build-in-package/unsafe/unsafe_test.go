package unsafe

import (
	"fmt"
	"sync/atomic"
	"testing"
	"unsafe"
)

type T struct {
	value int
}

func Swap(dest **T, old, new *T) bool {
	udest := (*unsafe.Pointer)(unsafe.Pointer(dest))
	return atomic.CompareAndSwapPointer(udest,
		unsafe.Pointer(old),
		unsafe.Pointer(new),
	)
}
func TestSwap(t *testing.T) {
	x := &T{42}
	n := &T{50}
	fmt.Println(*x, *n)

	p := x
	Swap(&x, p, n)
	fmt.Println(x, *p, *n)
}
