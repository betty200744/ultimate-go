package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
	ch := make(chan []byte)
	go func(ch chan []byte) {
		resp, _ := http.Get("https://example.com")
		defer resp.Body.Close()
		b, _ := ioutil.ReadAll(resp.Body)
		ch <- b
		close(ch)
	}(ch)
	fmt.Println("Program continues....")
	fmt.Println("Done, result is ", string(<-ch))
}
