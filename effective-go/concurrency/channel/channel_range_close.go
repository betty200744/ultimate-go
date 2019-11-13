package main

import "fmt"

func febonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // channel range and close
}

func main() {
	ch3 := make(chan int)
	go febonacci(10, ch3)
	fmt.Println(<-ch3, <-ch3, <-ch3) // 0, 1, 1  注意 <- 只取一次， 如果全取需要 range ch2 {}
	for i := range ch3 {             // 用range把ch3的value都取出来
		fmt.Println(i)
	}
}
