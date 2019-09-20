package main

import (
	"fmt"
)

func fibonacci(c, quit chan int) {
	for {
		select {
		case c <- 1: // 写入c channelΩ
		case <-quit: // 读取quit channel
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-c) // 读取channel
		}
		quit <- 0 // 写入quit channel
	}()
	fibonacci(c, quit)
}
