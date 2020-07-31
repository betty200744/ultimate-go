package main

import (
	"fmt"
	"runtime"
)

func main() {
	numCPU := runtime.NumCPU()
	fmt.Println("NumCPU: ", numCPU)
	numGoroutine := runtime.NumGoroutine()
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	fmt.Println("NumGoroutine: ", numGoroutine)
}
