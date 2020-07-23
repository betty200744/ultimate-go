package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNode_Add(t *testing.T) {
	tree := New(2)
	tree.Add(3)
	tree.Print()
}

func TestNode_Contains(t *testing.T) {
	tree := New(2)
	tree.Add(3)
	tree.Add(1)
	tree.Print()
	assert.Equal(t, tree.Contains(1), true)
	assert.Equal(t, tree.Contains(3), true)
	assert.Equal(t, tree.Contains(2), true)
	assert.Equal(t, tree.Contains(5), false)
}
