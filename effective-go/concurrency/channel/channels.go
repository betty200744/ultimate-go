package main

import (
	"context"
	"fmt"
	"time"
)

func NormalJobRun() {
	fmt.Println("enter NormalJobRun")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers
	go func(ctx context.Context) {
		fmt.Println("enter go func")

		for {
			select {
			case <-time.After(time.Second * 100):
				fmt.Println("timeout")
			case <-ctx.Done():
				fmt.Println("NormalJobRun return, so for loop return")
				return
			}
		}
	}(ctx)
	time.Sleep(time.Second * 3)
	fmt.Println("exit NormalJobRun")

	return
}
func main() {
	go NormalJobRun()
	select {}
}
