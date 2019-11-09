package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)
import "runtime"

// goroutine, a fun or method concurrently running in background
// goroutine, a function executing concurrently with other goroutines in the same address space.
// goroutine, call with go , then goroutine
// goroutine, not block the current program
// goroutine, scheduler concurrently,
// goroutine, main goroutine not scheduler spawns goroutine when done
// goroutine, main goroutine can manually scheduler spawns goroutine
// goroutine, when main goroutine sleep, then main
func main() {
	var result int
	f, _ := os.Create("trace.out")
	defer f.Close()
	processors := runtime.GOMAXPROCS(0)
	fmt.Println(processors)
	defer trace.Stop()

	for i := 0; i < processors; i++ {
		trace.Start(f)
		go func(i int) {
			fmt.Println(i)
			time.Sleep(time.Second)
		}(i)
	}
	fmt.Println("result =", result)

	fmt.Println("goroutine", runtime.NumGoroutine())
	time.Sleep(time.Second * 3)
}
