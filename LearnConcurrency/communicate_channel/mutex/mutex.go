package main

import (
	"fmt"
	"sync"
)

var (
	g int
)

func Add(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock() // acquire lock
	g = g + 1
	m.Unlock() // release lock
	wg.Done()
}
func main() {
	// Mutex, before operator, mutex.Lock()
	// Mutex, after operator, mutex.Unlock()
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go Add(wg, mutex)
	}
	// waitgroup wait all goroutine
	wg.Wait()
	fmt.Println("value of i after 1000 operations is", g)
}
