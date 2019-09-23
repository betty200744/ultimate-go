package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	fmt.Println("timer1 is expired", <-timer1.C)

	timer2 := time.NewTimer(time.Second * 5)
	go func() {
		fmt.Println("timer2 is expired", <-timer2.C)
	}()
	isStop := timer2.Stop()
	if isStop {
		fmt.Println("timer2 is stop!")
	}
}
