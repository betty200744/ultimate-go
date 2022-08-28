package __findRepeatNumber

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestFindRepeatNumber(t *testing.T) {
	arr := []int{2, 3, 1, 2, 5, 3}
	res := FindRepeatNumber(arr)
	assert.Equal(t, res, 2)
}
