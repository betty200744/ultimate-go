package sync

import (
	"fmt"
	"sync"
	"testing"
)

var (
	state = make(map[int]int, 5)
	mutex = &sync.Mutex{} // 读写互斥， 写的时候不让读
)

func stateMutex() {
	mutex.Lock()
	defer mutex.Unlock()
	go func() {
		state[1] = 1
		state[2] = 2
	}()
}

func Test_Mutex(t *testing.T) {
	stateMutex()
	state[1] = 3
	fmt.Println(state)
}
