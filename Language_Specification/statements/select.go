package main

import "fmt"

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan int, 1)

	ch1 <- "hello"
	ch2 <- 1
	for i := 1; i < 3; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("this from ch1" + msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}
}
