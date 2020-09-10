package _70_maximumSwap

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestMaximumSwap(t *testing.T) {
	res := MaximumSwap(2736)
	assert.Equal(t, 7236, res)
}
func TestMaximumSwap2(t *testing.T) {
	res := MaximumSwap(9973)
	assert.Equal(t, 9973, res)
}
func TestMaximumSwap3(t *testing.T) {
	res := MaximumSwap(10)
	assert.Equal(t, 10, res)
}
func TestMaximumSwap4(t *testing.T) {
	res := MaximumSwap(20)
	assert.Equal(t, 20, res)
}
func TestMaximumSwap5(t *testing.T) {
	res := MaximumSwap(98368)
	assert.Equal(t, 98863, res)
}
