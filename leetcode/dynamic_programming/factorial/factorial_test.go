package factorial

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestFactorial(t *testing.T) {
	res := Factorial(4)
	assert.Equal(t, 24, res)
}
func TestFactorialDP(t *testing.T) {
	res := FactorialDP_Tabulation(4)
	assert.Equal(t, 24, res)
}

func TestFactorialDP_Memoization(t *testing.T) {
	res := FactorialDP_Memoization(4)
	assert.Equal(t, 24, res)
}
