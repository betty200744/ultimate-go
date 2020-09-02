package _43_diameterOfBinaryTree

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestHeight(t *testing.T) {
	root := NewTreeNode(1)
	root.Left = NewTreeNode(2)
	root.Left.Left = NewTreeNode(4)
	root.Left.Right = NewTreeNode(5)
	root.Right = NewTreeNode(3)
	res := Height(root)
	assert.Equal(t, 3, res)
}

func TestDiameterOfBinaryTree(t *testing.T) {
	root := NewTreeNode(1)
	root.Left = NewTreeNode(2)
	root.Left.Left = NewTreeNode(4)
	root.Left.Right = NewTreeNode(5)
	root.Right = NewTreeNode(3)
	res := DiameterOfBinaryTree(root)
	assert.Equal(t, 3, res)
}
