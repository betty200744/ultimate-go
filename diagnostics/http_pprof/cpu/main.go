package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func counter() {
	slice := make([]int, 0)
	c := 1
	for i := 0; i < 100000; i++ {
		c = i + 1 + 2 + 3 + 4 + 5
		slice = append(slice, c)
	}
}

func workForever() {
	for {
		go counter()
		time.Sleep(1 * time.Millisecond)
	}
}

func httpGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`hello http pprof`))
}

func main() {
	// fun loop use cpu
	go workForever()
	// just http service api, also pprof api for generator pprof file
	http.HandleFunc("/get", httpGet)
	fmt.Println("http service: localhost:5017")
	http.ListenAndServe("localhost:5017", nil)
}
