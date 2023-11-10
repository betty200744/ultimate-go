package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`hello, world!\n`))
	})
	http.ListenAndServe(":8080", nil)
}
