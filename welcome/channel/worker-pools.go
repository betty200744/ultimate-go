package main

import (
	"fmt"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker id:", id, "started job:", job)
		results <- job * 100
	}
}

func main() {
	jobs := make(chan int, 30)
	results := make(chan int, 30)

	for w := 1; w < 4; w++ {
		go worker(w, jobs, results)
	}

	for i := 1; i < 10; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 1; j < 10; j++ {
		fmt.Println(<-results)
	}

}
