package main

import (
	"fmt"
	"time"
)

func TickerTimeout(timeout time.Duration) {
	ticker := time.NewTicker(timeout)
	select {
	case <-ticker.C:
		fmt.Println("NewTicker timeout")
	}
}
func TimeAfterTimeout(timeout time.Duration) {
	select {
	case <-time.After(timeout):
		fmt.Println("time.After timeout")
	}
}
