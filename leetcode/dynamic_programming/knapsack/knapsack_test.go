package knapsack

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestKnapsack(t *testing.T) {
	res := Knapsack([]Item{{weight: 10, value: 60}, {weight: 20, value: 100}, {weight: 30, value: 120}}, 50)
	assert.Equal(t, 220, res)
}
