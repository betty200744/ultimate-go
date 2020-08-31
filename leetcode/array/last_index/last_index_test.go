package last_index

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLastIndex(t *testing.T) {
	nums := []int{1, 7, 3, 6, 5, 6}
	res := LastIndex(nums, 6)
	assert.Equal(t, 5, res)
}
