package time

import (
	"fmt"
	"time"
)

func TimeDuration() {
	// time.Duration
	now := time.Now()
	time.Sleep(time.Second)
	duration := time.Since(now)
	fmt.Println("this is Duration.String, Seconds:", duration.String())
	fmt.Println("this is Duration.Hours:", duration.Hours())
	fmt.Println("this is Duration.Minutes:", duration.Minutes())
	fmt.Println("this is Duration.Seconds:", duration.Seconds())
}
