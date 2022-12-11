package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

// cpu performance
func counter(wg *sync.WaitGroup) {
	wg.Done()
	slice := make([]int, 0)
	c := 1
	for i := 0; i < 100000; i++ {
		c = i + 1 + 2 + 3 + 4 + 5
		slice = append(slice, c)
	}
}
func counter1(wg *sync.WaitGroup) {
	wg.Done()
	slice := make([]int, 100000)
	c := 1
	for i := 0; i < 100000; i++ {
		c = i + 1 + 2 + 3 + 4 + 5
		slice[i] = c
	}
}
func counter2(wg *sync.WaitGroup) {
	wg.Done()
	slice := [100000]int{}
	c := 1
	for i := 0; i < 100000; i++ {
		c = i + 1 + 2 + 3 + 4 + 5
		//slice = append(slice, c)
		slice[i] = c
	}
}

// go run trace.go
// go tool trace trace.pprof
func main() {
	runtime.GOMAXPROCS(2)
	var traceProfile = flag.String("traceprofile", "trace.pprof", "write trace profile to file")
	flag.Parse()

	if *traceProfile != "" {
		f, err := os.Create(*traceProfile)
		if err != nil {
			log.Fatal(err)
		}
		trace.Start(f)
		defer f.Close()
		defer trace.Stop()
	}

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go counter2(&wg)
	}
	wg.Wait()
}
