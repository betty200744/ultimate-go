package _4_longestCommonPrefix

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	res := LongestCommonPrefix([]string{"flower", "flow", "flight"})
	assert.Equal(t, "fl", res)
	res = LongestCommonPrefix([]string{"dog", "racecar", "car"})
	assert.Equal(t, "", res)
	res = LongestCommonPrefix([]string{"a"})
	assert.Equal(t, "a", res)
	res = LongestCommonPrefix([]string{"aaa", "aa", "aaa"})
	assert.Equal(t, "aa", res)
	res = LongestCommonPrefix([]string{"aca", "cba"})
	assert.Equal(t, "", res)
}
