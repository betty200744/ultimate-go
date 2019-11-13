package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// channel, 概念, goroutines communicate, use channel transport data between goroutine
// channel, 概念, goroutines communicate, one A read , one B write
// channel, 关键字, chan
// channel, 定义, declaring var, channel is nil, cannot read and write
// channel, 类型, data write: ch <- data , data read: i <- ch
// channel, 分配 or 创建, allocated make, channel, unbuffered, capacity 0 ,channel is address, can read and write
// channel, 分配 or 创建, allocated make, channel, buffered, capacity 0, full then deadlock

// channel, 读写方式, chan<- 写,  <-chan读
// channel, 读写方式, for, loop read channel,need break loop, 一般不用
// channel, 读写方式, for range, loop read channel, will break when chan is close, 常用, 单chan遍历读写
// channel, 读写方式, select, (case, timeout, default): once read channel,   val := <-ch5, 常用, 多chan读写, 部分成功就行
// channel, 读写方式, WaitGroup, read all channel, WaitGroup, wg.add, wg.done, wg.wait,常用,多chan读写, 全部成功才行
// channel, 读写方式, worker pool, 一般两chan， 一个tasks, 一个results，producer写tasks,  worker读tasks写results, consumers读results， 多chan遍历读写

// channel, 关闭, closing channel, when no more data te sent
// channel, 判断, present channel, val, ok := <-channel , 判断channel is closed?

var (
	g int
)

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

func getUrl(wg *sync.WaitGroup, url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("waitGroup: read all channel, WaitGroup done: ", res.Status)
	wg.Done()
}

func Add(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock() // acquire lock
	g = g + 1
	m.Unlock() // release lock
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
	// default is faster then ch5
	ch5 := make(chan int)
	go squares(ch5)
	select {
	case val := <-ch5:
		fmt.Println("select: case and default read ch5: ", val)
	default:
		fmt.Println("select: default")
	}
	// ch5 has val, so run case
	select {
	case val := <-ch5:
		fmt.Println("select: case and timeout read ch5:", val)
	case <-time.After(time.Millisecond * 20):
		fmt.Println("select: timeout")
	}
	// ch6 is nil, so run default
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

	// Mutex, before operator, m.Lock()
	// Mutex, after operator, m.Unlock()
	wg2 := &sync.WaitGroup{}
	m := &sync.Mutex{}
	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go Add(wg2, m)
	}
	wg2.Wait()
	//
	m2 := &sync.Mutex{}
	for i := 0; i < 1000; i++ {
		go func() {
			m2.Lock()
			g = g + 1
			m2.Unlock()
		}()
	}

	time.Sleep(time.Second * 2)
	fmt.Println("value of i after 1000 operations is", g)
}
