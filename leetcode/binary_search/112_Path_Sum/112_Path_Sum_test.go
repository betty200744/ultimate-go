package HasPathSum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	tree := NewNode(5)
	tree.Left = NewNode(4)
	tree.Right = NewNode(8)
	tree.Left.Left = NewNode(11)
	tree.Right.Left = NewNode(13)
	tree.Right.Right = NewNode(4)
	tree.Left.Left.Left = NewNode(7)
	tree.Left.Left.Right = NewNode(2)
	tree.Right.Right.Right = NewNode(1)
	BreadthFirst(tree)
	res := HasPathSum(tree, 22)
	assert.Equal(t, true, res)
	res = HasPathSum(tree, 18)
	assert.Equal(t, true, res)
}
