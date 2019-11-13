package main

import (
	"fmt"
	"time"
)

func main() {
	a := 1
	go func() {
		a = 2
		fmt.Println(a)
	}()
	go func() {
		a = 3
		fmt.Println(a)
	}()
	time.Sleep(time.Second)
	fmt.Println(a)
}
