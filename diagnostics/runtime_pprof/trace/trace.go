package main

import (
	"flag"
	"log"
	"os"
	"runtime/trace"
	"sync"
)

func counter(wg *sync.WaitGroup) {
	wg.Done()
	slice := []int{0}
	c := 1
	for i := 0; i < 100000; i++ {
		c = i + 1 + 2 + 3 + 4 + 5
		slice = append(slice, c)
	}
}

func main() {
	//runtime.GOMAXPROCS(1)
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
		go counter(&wg)
	}
	wg.Wait()
}
