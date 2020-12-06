package time

import (
	"fmt"
	"time"
)

func TimeOutByTimer() {
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

func TimeOutByChannel() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("time out 1")
	}
}
