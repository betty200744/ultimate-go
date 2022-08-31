package index_range

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndexRange(t *testing.T) {
	nums := []int{1, 7, 3, 6, 5, 6, 7}
	res := IndexRange(nums, 6)
	assert.Equal(t, []int{3, 5}, res)
	res = IndexRange(nums, 7)
	assert.Equal(t, []int{1, 6}, res)
}
