package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	a := make(chan int, 9)
	for i := 0; i < 9; i++ {
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
	time.Sleep(time.Second * 3)
	cancel()
	time.Sleep(time.Second * 6)
	/*	sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
		select {
		case <-sigterm:
			log.Println("terminating: via signal")

		}*/

}
