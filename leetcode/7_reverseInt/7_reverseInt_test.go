package __reverseInt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseInt(t *testing.T) {
	res := ReverseInt(12345)
	assert.Equal(t, 54321, res)
}

func TestReverseInt2(t *testing.T) {
	res := ReverseInt(1534236469)
	assert.Equal(t, 0, res)
}
func TestReverseInt3(t *testing.T) {
	res := ReverseInt(-12345)
	assert.Equal(t, -54321, res)
}
func TestReverseInt4(t *testing.T) {
	res := ReverseInt(900000)
	assert.Equal(t, 9, res)
}
