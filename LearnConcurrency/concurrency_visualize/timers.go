package main

import (
	"fmt"
	"time"
)

func timer() <-chan int {
	c := make(chan int)
	go func() {
		time.Sleep(time.Second)
		c <- 1
	}()
	return c
}

func RunTimer() {
	for i := 0; i < 24; i++ {
		r := timer()
		fmt.Println(<-r)
	}
}
