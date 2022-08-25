package GCD

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestGCD(t *testing.T) {
	assert.Equal(t, 17, GCD(17, 34))
	assert.Equal(t, 34, GCD(0, 34))
	assert.Equal(t, 17, GCD(17, 0))
	assert.Equal(t, 1, GCD(49, 50))
	assert.Equal(t, 1, GCD(50, 49))
	assert.Equal(t, 1, GCD(50, 17))
}
