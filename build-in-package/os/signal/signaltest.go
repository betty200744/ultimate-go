package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	c1 := make(chan os.Signal, 1)
	go func() {
		signal.Notify(c1, os.Interrupt, os.Kill)
	}()

	c2 := make(chan os.Signal, 1)
	signal.Notify(c2, os.Interrupt, os.Kill)
	fmt.Println("got signal 1", <-c1)
	fmt.Println("got signal 2", <-c2)
	time.Sleep(time.Second)
}
