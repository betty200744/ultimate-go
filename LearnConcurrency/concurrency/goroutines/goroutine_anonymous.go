package main

import (
	"fmt"
	"time"
)

// foo goroutine
func foo(a int) {
	a = 2
}

// main goroutine
func main() {
	a := 1
	go foo(a)
	fmt.Println("after foo goroutine: ", a)
	// anonymous goroutine, will change a, a is in closures
	go func() {
		a = 2
	}()
	time.Sleep(time.Second)
	fmt.Println("after anonymous goroutine: ", a)
}
