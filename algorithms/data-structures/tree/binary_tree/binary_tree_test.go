package binary_tree

import (
	"fmt"
	"testing"
)

func CreateBinaryTree() *Node {
	tree := NewNode(1)
	tree.Left = NewNode(2)
	tree.Right = NewNode(3)
	tree.Left.Left = NewNode(4)
	tree.Left.Right = NewNode(5)
	return tree
}

func TestBreadthFirst(t *testing.T) {
	tree := CreateBinaryTree()
	BreadthFirst(tree)
	fmt.Println("")
}

func TestPreOrder(t *testing.T) {
	tree := CreateBinaryTree()
	PreOrder(tree)
	fmt.Println("")

}
func TestInOrder(t *testing.T) {
	tree := CreateBinaryTree()
	InOrder(tree)
	fmt.Println("")
}
func TestPostOrder(t *testing.T) {
	tree := CreateBinaryTree()
	PostOrder(tree)
	fmt.Println("")
}
