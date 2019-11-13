package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

// goroutine, 概念：it is a function executing concurrently with other goroutines in the same address space
// goroutine, 关键字: go , then goroutine
// goroutine, 通信: channel, 详情见channel-cheat-sheet.go
// goroutine, 传参：闭合的, 所以需要传入参数
// goroutine, 传参：for变量, 需要传入i到go fn
// goroutine, 传参：chan变量,需要传入chan到go fn
// goroutine, 传参：如果fn写ch, 那么形参: chan<-
// goroutine, 传参：如果fn读ch, 那么形参: <-chan
// goroutine, 读写全局变量, 需要Mutex, m.Lock(), m.Unlock(), 否则即使传pointer也不能保证原子性
// goroutine, anonymous goroutine,
// Processes, 进程， 对应多个Threads, 多个Ports
// Threads, 线程,

var ggg int

func main() {
	c := make(chan int)     //通信
	go func(c chan<- int) { // 形参
		list := []int{3, 2, 1}
		sort.Ints(list)
		c <- 1
	}(c) // 传参
	time.Sleep(time.Second)
	go func() {
		list := []int{3, 2, 1}
		sort.Ints(list)
		c <- 2 // 直接读写chan
	}()

	m := sync.Mutex{} // mutex读写全局变量
	for i := 0; i < 999; i++ {
		go func() {
			m.Lock()
			ggg = ggg + 1
			m.Unlock()
		}()
	}

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(ggg)
}
