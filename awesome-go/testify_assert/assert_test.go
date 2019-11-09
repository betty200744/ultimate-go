package testify_assert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {

	//assert.Equal(t, a, b, "The two words should be the same.")
	assert.Greater(t, 2, 1)
	assert.Greater(t, float64(2), float64(1))
	assert.Greater(t, "b", "a")
}
