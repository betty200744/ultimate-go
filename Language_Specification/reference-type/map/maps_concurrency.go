package main

import (
	"sync"
)

func main() {
	m := make(map[string]int, 1)
	m[`foo`] = 1

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 1000; i++ {
			m[`foo`]++
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			m[`foo`]++
		}
		wg.Done()
	}()
	wg.Wait()
}
