package binary_search_tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNode_Insert(t *testing.T) {
	tree := NewNode(2)
	tree.Insert(3)
}
func TestNode_Delete(t *testing.T) {
	tree := NewNode(1)
	tree.Insert(2)
	tree.Insert(3)
	BreadthFirst(tree)
	tree.Delete(3)
	fmt.Println("")
	BreadthFirst(tree)
}
func TestNode_Find(t *testing.T) {
	tree := NewNode(2)
	tree.Insert(3)
	tree.Insert(1)
	assert.Equal(t, tree.Find(1), true)
	assert.Equal(t, tree.Find(3), true)
	assert.Equal(t, tree.Find(2), true)
	assert.Equal(t, tree.Find(5), false)
}
func TestNode_FindMin(t *testing.T) {
	tree := NewNode(2)
	tree.Insert(3)
	tree.Insert(1)
	tree.FindMin()
}
func TestNode_FindMax(t *testing.T) {
	tree := NewNode(2)
	tree.Insert(3)
	tree.Insert(1)
	tree.FindMax()
}
