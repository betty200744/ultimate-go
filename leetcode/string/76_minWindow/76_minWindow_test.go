package _6_minWindow

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestIncludeAll(t *testing.T) {
	res := IncludeAll("BANC", "ABC")
	assert.Equal(t, true, res)
	res = IncludeAll("BECO", "ABC")
	assert.Equal(t, false, res)
	res = IncludeAll("aa", "a")
	assert.Equal(t, true, res)
	res = IncludeAll("a", "a")
	assert.Equal(t, true, res)
}
func TestMinWindow(t *testing.T) {
	res := MinWindow("ADOBECODEBANC", "ABC")
	assert.Equal(t, "BANC", res)
	res = MinWindow("aa", "a")
	assert.Equal(t, "a", res)
	res = MinWindow("a", "a")
	assert.Equal(t, "a", res)
	res = MinWindow("bbaa", "aba")
	assert.Equal(t, "baa", res)
}
func TestMinWindow2(t *testing.T) {
	res := MinWindow2("cabwefgewcwaefgcf", "cae")
	assert.Equal(t, "cwae", res)
}
