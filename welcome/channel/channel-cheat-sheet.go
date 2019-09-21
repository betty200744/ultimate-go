package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// channel, goroutines communicate, use channel transport data between goroutine
// channel, goroutines communicate, one A read , one B write
// channel, goroutines communicate, A block means wait B

// channel, declaring channel, channel is nil, cannot read and write
// channel, make channel, unbuffered, capacity 0 ,channel is address, can read and write
// channel, make channel, buffered, capacity 0, full then deadlock

// channel, data write: ch <- data , data read: i <- ch
// channel, receive only channel, revoch := make(<-chan int)
// channel, send only channel, sendoch := make(chan<- int)
// channel, channel operations,  are blocking by default
// channel, channel operations,  are tell the scheduler to schedule another goroutine
// channel, channel operations, 当只有main channel时,  此时main channel deadlock!
// channel, channel operations, 当只有spawn channel时, 此时spawn channel deadlock!

// channel, loop read channel, for loop, need break loop
// channel, loop read channel, for range, , will break when chan is close
// channel, once read channel, select case:  val := <-ch5
// channel, once read channel, select default: read channel failed
// channel, once read channel, select timeout: read channel failed
// channel, once read channel, select empty: deadlock

// channel, read all channel, WaitGroup: init wg , before goroutine wg.add, after goroutine wg.done, last main goroutine wg.wait
// channel, worker pool,
// channel, Mutex,
// channel, closing channel, when no more data te sent
// channel, present channel, val, ok := <-channel , 判断channel is closed?

// goroutine, main goroutine必须把channel通过参数传入spawn goroutine, 否则会有slide effect
// goroutine, spawn goroutine必须把数据通过channel传到main goroutine， 否则数据无法到main goroutine

func squares(ch3 chan int) {
	for i := 0; i < 9; i++ {
		ch3 <- i * i
	}
	close(ch3)
}

func workers(jobs chan int, results chan string) {
	for job := range jobs {
		results <- fmt.Sprintf("worker pools: job: %d, after worker: %d", job, job*job)
	}
	close(results)
}

func getUrl(wg *sync.WaitGroup, url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("waitGroup: read all channel, WaitGroup done: ", res.Status)
	wg.Done()
}

func main() {
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

	// for loop: loop read channel, need break loop
	ch3 := make(chan int)
	go squares(ch3)
	for {
		val, ok := <-ch3
		if ok {
			fmt.Println("for loop: read ch3: ", val)
		} else {
			fmt.Printf("for loop: read ch3: break for!, value: %d, ok: %t", val, ok)
			break // need a condition to break the loop
		}
	}
	// for range: loop read channel, no need break loop
	ch4 := make(chan int)
	go squares(ch4)
	for value := range ch4 {
		fmt.Println("for range: loop read ch4:", value)
	}

	// select
	ch5 := make(chan int)
	go squares(ch5)
	select {
	case val := <-ch5:
		fmt.Println("select: case and default read ch5: ", val)
	default:
		fmt.Println("select: default")
	}
	select {
	case val := <-ch5:
		fmt.Println("select: case and timeout read ch5:", val)
	case <-time.After(time.Millisecond * 20):
		fmt.Println("select: timeout")
	}
	var ch6 chan int
	select {
	case val := <-ch6:
		fmt.Println("select: case read nil ch6:", val)
	default:
		fmt.Println("select: case read nil ch6 default")
	}

	// WaitGroup: first: init wg
	// WaitGroup: then: before goroutine wg.add
	// WaitGroup: then: after goroutine wg.done
	// WaitGroup: last: main goroutine wg.wait
	wg := &sync.WaitGroup{}
	urls := []string{"http://www.baidu.com", "https://gorm.io"}
	for _, url := range urls {
		wg.Add(1)
		go getUrl(wg, url)
	}
	wg.Wait()

	// Worker pool
	jobs := make(chan int)
	results := make(chan string)
	go squares(jobs)
	go workers(jobs, results)
	for result := range results {
		fmt.Println(result)
	}

	// Mutex

}
