package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func fibRecursive(n int64) int64 {
	if n < 2 {
		return n
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

func fibIterative(n int64) int64 {
	if n < 2 {
		return n
	}

	var fib = int64(1)
	var prevFib = fib
	var tmp int64

	for i := int64(2); i < n; i++ {
		tmp = fib
		fib += prevFib
		prevFib = tmp
	}

	return fib
}

type fibResponse struct {
	Number    int64 `json:"number"`
	Fibonacci int64 `json:"fibonacci"`
}

func FibHandler(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.FormValue("n"))
	if err != nil {
		n = 1
	}

	t := r.FormValue("type")
	if t == "" {
		t = "recursive"
	}

	resp := fibResponse{Number: int64(n), Fibonacci: 0}

	switch t {
	case "recursive":
		resp.Fibonacci = fibRecursive(resp.Number)
	case "iterative":
		resp.Fibonacci = fibIterative(resp.Number)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
func GinFibHandler(c *gin.Context) {
	r := c.Request
	n, err := strconv.Atoi(r.FormValue("n"))
	if err != nil {
		n = 1
	}

	t := r.FormValue("type")
	if t == "" {
		t = "recursive"
	}

	resp := fibResponse{Number: int64(n), Fibonacci: 0}

	switch t {
	case "recursive":
		resp.Fibonacci = fibRecursive(resp.Number)
	case "iterative":
		resp.Fibonacci = fibIterative(resp.Number)
	}
	w := c.Writer
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
