package _26_invertTree

import (
	"fmt"
	"testing"
)

func TestInvertTree(t *testing.T) {
	root := NewNode(4)
	root.Left = NewNode(2)
	root.Left.Left = NewNode(1)
	root.Left.Right = NewNode(3)
	root.Right = NewNode(7)
	root.Right.Left = NewNode(6)
	root.Right.Right = NewNode(9)
	BreadthFirst(root)
	root = InvertTree(root)
	fmt.Println("")
	BreadthFirst(root)
}
