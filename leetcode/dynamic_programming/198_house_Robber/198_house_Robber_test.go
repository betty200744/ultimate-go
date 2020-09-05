package _98_house_Robber

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestRob(t *testing.T) {
	res := Rob([]int{1, 2, 3, 1})
	assert.Equal(t, 4, res)
	res = Rob([]int{2, 7, 9, 3, 1})
	assert.Equal(t, 12, res)
}
