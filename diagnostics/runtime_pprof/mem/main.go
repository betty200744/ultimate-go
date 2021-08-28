package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"sync"
)

// len 100000, mem
func counter() {
	slice := [100000]int{0}
	c := 1
	for i := 0; i < 100000; i++ {
		c = i + 1 + 2 + 3 + 4 + 5
		slice[i] = c
	}
}

// generator pprof: go run main.go
// analyze pprof sample data: go tool pprof mem.pprof
// go tool pprof mem.pprof
func main() {
	var memProfile = flag.String("memprofile", "mem.pprof", "write mem profile to file")
	flag.Parse()

	// 100 counter gorouting
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go counter()
		wg.Done()
	}
	wg.Wait()

	//profiling memory status
	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not start memory profile")
		}
	}
}
