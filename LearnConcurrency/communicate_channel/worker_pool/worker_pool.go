package main

import "fmt"

func squares(ch3 chan int) {
	for i := 0; i < 9; i++ {
		ch3 <- i * i
	}
	close(ch3) // must close
}

func workers(jobs chan int, results chan string) {
	for job := range jobs {
		results <- fmt.Sprintf("worker pools: job: %d, after worker: %d", job, job*job)
	}
	close(results)
}

func main() {
	// Worker pool, one write,  jobs write chan, then close chan
	// Worker pool, one read, workers read chan and write results chan then close chan
	// Worker pool, then range results chan
	jobs := make(chan int)
	results := make(chan string)
	go squares(jobs)
	go workers(jobs, results)
	for result := range results {
		fmt.Println(result)
	}
}
