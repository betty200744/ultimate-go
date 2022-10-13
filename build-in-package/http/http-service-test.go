package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type H struct {
}

func (h *H) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	var body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	// And now set a new body, which will simulate the same data we read:
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	fmt.Println("body: ", string(body))
	fmt.Fprint(w, "/hello")
}

func main() {
	http.HandleFunc("/hello", Hello)
	http.ListenAndServe("localhost:3004", nil)
}
