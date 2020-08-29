package channel

import (
	"context"
	"fmt"
	"time"
)

// for loop, exit when i >3
func Channel1() {
	queue := make(chan int, 3)
	for i := 0; i < 3; i++ {
		queue <- i
	}
	for i := 0; i < 3; i++ {
		fmt.Println(<-queue)
	}
}

// exit for range loop by close channel
func Channel2() {
	queue := make(chan int, 3)
	for i := 0; i < 3; i++ {
		queue <- i
	}
	// need close when use range, 即exit the for loop, 避免deadlock
	close(queue)
	for ele := range queue {
		fmt.Println(ele)
	}
}

// exit for while loop by ctx.Done
func Channel3() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func(ctx context.Context) {
		for {
			select {
			case <-time.After(time.Second * 100):
			// exit the select loop
			case <-ctx.Done():
				fmt.Println("NormalJobRun return, so exit for loop ")
				return
			}
		}
	}(ctx)
}

// receive only channel 不能 close，直接编译错误。但 send only channel 是可以被正常 close 的
func Channel4() {
	sendOnly := make(chan<- int, 3)
	for i := 0; i < 3; i++ {
		sendOnly <- i
	}
	close(sendOnly)

	receiveOnly := make(<-chan int, 3)
	for i := 0; i < 3; i++ {
		fmt.Println(<-receiveOnly)
	}
	//close(receiveOnly) error
}
