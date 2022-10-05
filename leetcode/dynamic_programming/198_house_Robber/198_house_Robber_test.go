package _98_house_Robber

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestRob(t *testing.T) {
	res := Rob([]int{1, 2, 3, 4, 5, 6, 7})
	res = Rob([]int{1, 2, 3, 1})
	assert.Equal(t, 4, res)
	res = Rob([]int{2, 7, 9, 3, 1})
	assert.Equal(t, 12, res)
	res = Rob([]int{1, 3, 1, 3, 100})
	res = Rob([]int{4, 1, 1, 5})
	assert.Equal(t, 9, res)
}
func TestRobDP(t *testing.T) {
	res := RobDP([]int{1, 2, 3, 4, 5, 6, 7})
	res = RobDP([]int{1, 2, 3, 1})
	assert.Equal(t, 4, res)
	res = RobDP([]int{2, 7, 9, 3, 1})
	assert.Equal(t, 12, res)
	res = RobDP([]int{1, 3, 1, 3, 100})
	res = RobDP([]int{4, 1, 1, 5})
	assert.Equal(t, 9, res)
}
