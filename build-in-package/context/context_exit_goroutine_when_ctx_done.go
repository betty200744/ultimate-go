package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	a := make(chan int, 10)
	for i := 0; i < 10; i++ {
		a <- i
	}
	close(a)
	go func() {
		for {
			select {
			case v := <-a:
				{
					time.Sleep(time.Millisecond * 200)
					fmt.Println(v)
				}
			case <-ctx.Done():
				{
					log.Println("terminating: context cancelled")
					return

				}
			}
		}
	}()
	time.Sleep(time.Second * 2)
	cancel()
	time.Sleep(time.Second * 1)
}
