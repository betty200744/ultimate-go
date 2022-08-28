package KthSmallest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNode_KthSmallest(t *testing.T) {
	tree := NewNode(4)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	assert.Equal(t, 1, KthSmallest(tree, 1))
	assert.Equal(t, 4, KthSmallest(tree, 4))
	assert.Equal(t, 6, KthSmallest(tree, 6))
}
