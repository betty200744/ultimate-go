package _44_145_traversal

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestPreOrder(t *testing.T) {
	root := NewTreeNode(1)
	root.Right = NewTreeNode(2)
	root.Right.Left = NewTreeNode(3)
	res := PreOrder(root)
	assert.Equal(t, []int{1, 2, 3}, res)
}
func TestPreOrder2(t *testing.T) {
	root := NewTreeNode(1)
	root.Right = NewTreeNode(2)
	root.Right.Left = NewTreeNode(3)
	res := PreOrderIterative(root)
	assert.Equal(t, []int{1, 2, 3}, res)
}
func TestPostOrder(t *testing.T) {
	root := NewTreeNode(1)
	root.Right = NewTreeNode(2)
	root.Right.Left = NewTreeNode(3)
	res := PostOrder(root)
	assert.Equal(t, []int{3, 2, 1}, res)
}
func TestInOrder(t *testing.T) {
	root := NewTreeNode(1)
	root.Right = NewTreeNode(2)
	root.Right.Left = NewTreeNode(3)
	res := InOrder(root)
	assert.Equal(t, []int{1, 3, 2}, res)
}
