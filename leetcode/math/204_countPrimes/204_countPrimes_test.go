package _04_countPrimes

import (
	"fmt"
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestIsPrimes(t *testing.T) {
	assert.Equal(t, false, IsPrimes(1))
}
func TestCountPrimes(t *testing.T) {
	res := CountPrimes(10)
	assert.Equal(t, 3, res)
	res = CountPrimes(2)
	assert.Equal(t, 0, res)
}
func TestGenePrimes(t *testing.T) {
	res := GenePrimes(10)
	fmt.Println(res)
}
func TestGapPrimes(t *testing.T) {
	arr := GenePrimes(100)
	res := GapPrimes(arr, 3)
	fmt.Println(res)
}
