package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, mutex *sync.Mutex) {
	// goroutine 里面用的主线程变量，此时需要lock, 即lock 需要操作的变量x
	// mutex.Lock， 因为用到全局变量， 所以需要lock
	mutex.Lock()
	x = x + 1
	mutex.Unlock()
	wg.Done()
}

func main() {
	mutex := &sync.Mutex{}
	wait := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wait.Add(1)
		go increment(wait, mutex)
	}
	wait.Wait()
	fmt.Println(x)
}
