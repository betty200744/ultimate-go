package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	go func() {
		defer close(jobs)
		j, more := <-jobs
		if more {
			fmt.Println("received job", j)
		} else {
			fmt.Println("received all job")
		}
	}()

	for i := 1; i < 4; i++ {
		jobs <- i
		fmt.Println("send job", i)
	}
}
