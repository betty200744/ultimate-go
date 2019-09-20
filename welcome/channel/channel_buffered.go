package main

import "fmt"

func main() {
	// buffered channel
	ch1 := make(chan int, 1)
	ch2 := make(chan string, 2)
	ch1 <- 1233
	ch2 <- "abcd"
	fmt.Println(ch1)   // ch1 指向地址
	fmt.Println(ch2)   // ch2 指向地址
	fmt.Println(<-ch1) // 取值
	fmt.Println(<-ch2) // 取值
}
