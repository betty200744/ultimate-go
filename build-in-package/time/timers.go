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

func TimeoutByRetry(ttl time.Duration, deadline time.Time, maxRetry int64, fn func(ttl time.Duration) bool) {
	//ttl := time.Second * 3
	//deadline := time.Now().Add(time.Second * 7)
	//var maxRetry int64 = 2
	var timer *time.Timer
	for {
		// timeout
		if !time.Now().Before(deadline) {
			fmt.Println("total duration timeout")
			return
		}
		if fn(ttl) {
			fmt.Println("fn success!")
			return
		}
		// first
		if timer == nil {
			timer = time.NewTimer(time.Duration(int64(ttl) / maxRetry))
		} else {
			// retry
			timer.Reset(ttl)
		}
		// max retry time
		if maxRetry > 1 {
			maxRetry--
		} else {
			fmt.Println("max retry end")
			return
		}
		select {
		case <-timer.C:
			fmt.Println(time.Now())
		}
	}
}
