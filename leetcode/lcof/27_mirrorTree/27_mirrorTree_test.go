package _7_mirrorTree

import (
	"fmt"
	"testing"
)

func TestMirrorTree(t *testing.T) {
	tree := NewNode(4)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(6)
	tree.Insert(9)
	BreadthFirst(tree)
	MirrorTree(tree)
	fmt.Println("")
	BreadthFirst(tree)
}
