package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println("for: ", i)
		go func() {
			fmt.Println("go", i)
		}()
	}
	time.Sleep(time.Second)
}
