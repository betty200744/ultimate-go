package handler

import (
	"net/http/httptest"
	"testing"
)

func benchmarkFibRecursive(i int64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibRecursive(i)
	}
}

func BenchmarkFibRecursive1(b *testing.B)  { benchmarkFibRecursive(1, b) }
func BenchmarkFibRecursive2(b *testing.B)  { benchmarkFibRecursive(2, b) }
func BenchmarkFibRecursive3(b *testing.B)  { benchmarkFibRecursive(3, b) }
func BenchmarkFibRecursive10(b *testing.B) { benchmarkFibRecursive(10, b) }
func BenchmarkFibRecursive20(b *testing.B) { benchmarkFibRecursive(20, b) }
func BenchmarkFibRecursive40(b *testing.B) { benchmarkFibRecursive(40, b) }

func BenchmarkFibRecursiveHandler(b *testing.B) {
	r := httptest.NewRequest("GET", "/fib/?n=5&type=recursive", nil)
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		FibHandler(w, r)
	}
}

func benchmarkFibIterative(i int64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibIterative(i)
	}
}

func BenchmarkFibIterative1(b *testing.B)  { benchmarkFibIterative(1, b) }
func BenchmarkFibIterative2(b *testing.B)  { benchmarkFibIterative(2, b) }
func BenchmarkFibIterative3(b *testing.B)  { benchmarkFibIterative(3, b) }
func BenchmarkFibIterative10(b *testing.B) { benchmarkFibIterative(10, b) }
func BenchmarkFibIterative20(b *testing.B) { benchmarkFibIterative(20, b) }
func BenchmarkFibIterative40(b *testing.B) { benchmarkFibIterative(40, b) }

func BenchmarkFibIterativeHandler(b *testing.B) {
	r := httptest.NewRequest("GET", "/fib/?n=5&type=iterative", nil)
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		FibHandler(w, r)
	}
}
