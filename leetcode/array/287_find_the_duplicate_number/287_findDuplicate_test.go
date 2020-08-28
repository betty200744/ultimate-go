package _87_find_the_duplicate_number

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	arr := []int{1, 3, 4, 2, 2}
	res := FindDuplicate(arr)
	assert.Equal(t, 2, res)
}
