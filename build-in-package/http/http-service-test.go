package main

import (
	"fmt"
	"net/http"
)

type H struct {
}

func (h *H) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "/hello")
}

func main() {
	http.HandleFunc("/hello", Hello)
	http.ListenAndServe("localhost:3004", nil)
}
