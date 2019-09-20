package main

import (
	"fmt"
	"time"
)

var (
	start time.Time
)

func init() {
	start = time.Now()
}

func getChars(s string) {
	for _, value := range s {
		fmt.Printf("%c , time: %v \n", value, time.Since(start))
		time.Sleep(10 * time.Millisecond)
	}
}

func getDigits(d []int) {
	for _, value := range d {
		fmt.Printf("%d, time: %v \n", value, time.Since(start))
		time.Sleep(30 * time.Millisecond)
	}
}

func main() {
	go getChars("hello")
	go getDigits([]int{1, 2, 3, 4, 5})
	time.Sleep(time.Second)
	fmt.Printf("main end, time: %v", time.Since(start))
}
