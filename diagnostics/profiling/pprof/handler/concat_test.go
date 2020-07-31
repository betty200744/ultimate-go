package handler

import (
	"net/http/httptest"
	"testing"
)

func BenchmarkConcatHandler(b *testing.B) {
	r := httptest.NewRequest("GET", "/concat/?str=test&count=50", nil)
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		ConcatHandler(w, r)
	}
}

func benchmarkConcat(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		concat("test", n)
	}
}

func BenchmarkConcat10(b *testing.B) {
	benchmarkConcat(b, 10)
}

func BenchmarkConcat50(b *testing.B) {
	benchmarkConcat(b, 50)
}
