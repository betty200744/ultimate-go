package _04_countPrimes

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestIsPrimes(t *testing.T) {
	assert.Equal(t, true, IsPrimes(1))
}
func TestCountPrimes(t *testing.T) {
	res := CountPrimes(10)
	assert.Equal(t, 4, res)
	res = CountPrimes(2)
	assert.Equal(t, 0, res)
}
