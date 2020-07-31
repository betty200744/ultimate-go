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
