package sync

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func TestLoadUint64(t *testing.T) {
	var ops uint64
	for i := 0; i < 50; i++ {
		go func() {
			atomic.AddUint64(&ops, 1)
			time.Sleep(time.Millisecond)
		}()
	}

	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println(opsFinal)
}
