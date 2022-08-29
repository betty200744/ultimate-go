package _99_Binary_Tree_Right_Side_View

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

func TestRightSideView(t *testing.T) {
	tree := NewNode(1)
	tree.Left = NewNode(2)
	tree.Right = NewNode(3)
	tree.Left.Right = NewNode(5)
	tree.Right.Right = NewNode(4)
	BreadthFirst(tree)
	res := RightSideView(tree)
	fmt.Println("\n", res)
}

func TestLeftSideView(t *testing.T) {
	tree := NewNode(1)
	tree.Left = NewNode(2)
	tree.Right = NewNode(3)
	tree.Left.Left = NewNode(5)
	tree.Right.Left = NewNode(4)
	BreadthFirst(tree)
	res := LeftSideView(tree)
	fmt.Println("\n", res)
}
