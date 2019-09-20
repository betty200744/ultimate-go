package main

import (
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe("localhost:3005", http.FileServer(http.Dir("./")))
	if err != nil {
		log.Fatal(err)
	}
}
