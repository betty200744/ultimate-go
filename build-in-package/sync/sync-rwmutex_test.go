package sync

import (
	"container/list"
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	rwLock   sync.RWMutex
	connList list.List
)

func Test_RWMutex(t *testing.T) {
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	// Lock1
	rwLock.Lock()
	go func() {
		l.PushBack(3)
	}()
	rwLock.Unlock()
	// Lock2
	rwLock.Lock()
	go func() {
		l.PushBack(4)
	}()
	rwLock.Unlock()
	// RLock
	rwLock.RLock()
	go func() {
		fmt.Println(l.Front().Value)
		fmt.Println(l.Back().Value)
	}()
	rwLock.RUnlock()
	time.Sleep(time.Second)
	fmt.Println(l.Front().Value)
	fmt.Println(l.Back().Value)
}
