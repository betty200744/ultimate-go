package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool, 1)
	go func(ch chan bool) {
		fmt.Println("working....")
		time.Sleep(time.Second)
		fmt.Println("done")
		ch <- true
	}(done)

	fmt.Println(<-done) // wait ch done
	fmt.Println("main done")
}
