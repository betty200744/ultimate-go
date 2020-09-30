package _60_subarraySum

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestSubarraySum(t *testing.T) {
	res := SubarraySum([]int{1, 1, 1}, 2)
	assert.Equal(t, 2, res)
}
func TestSubarraySum2(t *testing.T) {
	res := SubarraySum([]int{1, 2, 1, 2, 1}, 3)
	assert.Equal(t, 2, res)
}

func TestSubarraySum3(t *testing.T) {
	res := SubarraySum([]int{1}, 1)
	assert.Equal(t, 1, res)
}
func TestSubarraySum4(t *testing.T) {
	res := SubarraySum([]int{1}, 0)
	assert.Equal(t, 0, res)
}
func TestSubarraySum5(t *testing.T) {
	res := SubarraySum([]int{-1, -1, 1}, 0)
	assert.Equal(t, 0, res)
}
