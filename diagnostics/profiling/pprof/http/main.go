package main

import (
	"fmt"
	"gobyexample/diagnostics/profiling/pprof/handler"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func foo() {
	select {}
}
func main() {
	go foo()
	go foo()
	go foo()
	go foo()
	http.HandleFunc("/concat/", handler.ConcatHandler)
	http.HandleFunc("/fib/", handler.FibHandler)
	log.Fatal(http.ListenAndServe(":6060", nil))
	fmt.Println(runtime.NumGoroutine())
}
