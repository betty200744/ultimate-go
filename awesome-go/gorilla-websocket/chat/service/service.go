package main

import (
	"net/http"
)

func main() {
	hub := NewHub()
	go hub.run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		WSHandle(hub, w, r)
	})
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic(err)
	}
}
