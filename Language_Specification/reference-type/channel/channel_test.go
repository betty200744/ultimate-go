package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	// declaring a channel
	var ch0 chan int
	ch1 := make(chan int)
	fmt.Println("declaring chan is :", ch0)
	fmt.Println("make chan is address: ", ch1)

	// use channel transport data between goroutine
	ch2 := make(chan int)
	go func(ch2 chan int) {
		ch2 <- 3
	}(ch2)
	fmt.Println("read ch2: ", <-ch2)

	// channel as the data type of channel
	cch7 := make(chan chan string)
	go func(cc chan chan string) {
		c := make(chan string)
		cc <- c
	}(cch7)
	ch7 := <-cch7
	go func(c chan string) {
		fmt.Println("hello, " + <-c)
	}(ch7)
	ch7 <- "betty"
	// for loop, exit when i >3
	Channel1()
	// exit for range loop by close channel
	Channel2()
	// exit for while loop by ctx.Done
	Channel3()
	// wait Channel3 goroutine
	time.Sleep(time.Second * 1)

}
