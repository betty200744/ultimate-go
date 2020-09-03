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

func TestFibonacciNumberDP(t *testing.T) {
	res := FibonacciNumberDP(2)
	assert.Equal(t, 1, res)
	res = FibonacciNumberDP(3)
	assert.Equal(t, 2, res)
	res = FibonacciNumberDP(4)
	assert.Equal(t, 3, res)
}
