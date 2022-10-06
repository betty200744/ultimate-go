package _09_fibonacciNumber

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestFibonacciNumber(t *testing.T) {
	res := FibonacciNumber(2)
	assert.Equal(t, 1, res)
	res = FibonacciNumber(3)
	assert.Equal(t, 2, res)
	res = FibonacciNumber(4)
	assert.Equal(t, 3, res)
}

func TestFib_DP_Tabulation(t *testing.T) {
	res := Fib_DP_Tabulation(2)
	assert.Equal(t, 1, res)
	res = Fib_DP_Tabulation(3)
	assert.Equal(t, 2, res)
	res = Fib_DP_Tabulation(4)
	assert.Equal(t, 3, res)
}
func TestFib_DP_memoized(t *testing.T) {
	res := Fib_DP_memoized(2)
	assert.Equal(t, 1, res)
	res = Fib_DP_memoized(3)
	assert.Equal(t, 2, res)
	res = Fib_DP_memoized(4)
	assert.Equal(t, 3, res)
}
