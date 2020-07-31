package main

import (
	"fmt"
)

func foo(m string, ch chan string) {
	ch <- m
}

func main() {
	ch2 := make(chan string) // 注意ch2先建， ch2先有value也就先输出
	ch1 := make(chan string)
	go foo("1", ch1)
	go foo("2", ch2)

	for i := 0; i < 2; i++ {
		select { // select , let you wait multi channel, 几个channel ， 则for loop 几次
		case msg1 := <-ch1:
			fmt.Println("this is msg" + msg1)
		case msg2 := <-ch2:
			fmt.Println("this is msg" + msg2)
		}
	}
}
