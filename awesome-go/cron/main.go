package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	_, _ = c.AddFunc("* * * * *", f("jobs hello", hello))
	c.Start()
	defer c.Stop()
	select {}
}

func hello() {
	fmt.Println("hello")
}
func f(msg string, f func()) func() {
	return func() {
		fmt.Println(msg)
		f()
	}
}
