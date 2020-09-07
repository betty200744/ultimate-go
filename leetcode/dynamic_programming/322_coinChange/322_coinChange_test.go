package _22_coinChange

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestCoinChange(t *testing.T) {
	res := CoinChange([]int{1, 2, 5}, 11)
	assert.Equal(t, 3, res)
}
func TestCoinChange2(t *testing.T) {
	res := CoinChange([]int{2}, 3)
	assert.Equal(t, -1, res)
}
func TestCoinChange3(t *testing.T) {
	res := CoinChange([]int{2}, 3)
	assert.Equal(t, -1, res)
}
