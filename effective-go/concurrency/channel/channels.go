package main

import "fmt"

func main() {
	// 新建一个chan int数据类型
	// 向chan里面写数据
	// 读chan里面的数据
	ch1 := make(chan int, 3)
	fmt.Println(&ch1)
	ch1 <- 1
	fmt.Println(<-ch1)

	// 一般用于goroutine写chan , 主线程读取chan
	ch2 := make(chan string, 2)
	go func() {
		ch2 <- "this from goroutine"
		ch2 <- "this from goroutine 2"
		ch2 <- "this from goroutine 3"
	}()
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
}
