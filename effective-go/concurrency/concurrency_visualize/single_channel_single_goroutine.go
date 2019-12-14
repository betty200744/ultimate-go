package main

import "fmt"

// single channel, single goroutine, one write, one read
func Single() {
	ch1 := make(chan int, 1)
	go func() {
		ch1 <- 42
	}()
	fmt.Println(<-ch1)
}
