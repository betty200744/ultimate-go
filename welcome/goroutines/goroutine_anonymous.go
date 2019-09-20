package main

import "fmt"

func foo(f string) {
	for i := 0; i < 3; i++ {
		fmt.Println(f, ":", i)
	}
}

// main goroutine
// foo goroutine
// anonymous goroutine
func main() {
	// in main goroutine
	foo("main direct")
	// in foo goroutine ， 开启一个线程执行该func
	go foo("goroutine 1")

	// anonymous goroutine
	go func(f string) { fmt.Println(f) }("goroutine 2 ")
	// wait a stand input
	fmt.Scanln()

	fmt.Println("main done")
}
