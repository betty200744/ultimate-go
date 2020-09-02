package _02_107_levelOrder

import (
	"fmt"
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestHeight(t *testing.T) {
	root := NewNode(3)
	root.Left = NewNode(9)
	root.Right = NewNode(20)
	root.Right.Left = NewNode(15)
	root.Right.Right = NewNode(7)
	res := Height(root)
	assert.Equal(t, 3, res)
}
func TestGivenLevelOrder(t *testing.T) {
	root := NewNode(3)
	root.Left = NewNode(9)
	root.Left.Left = NewNode(8)
	root.Left.Right = NewNode(10)
	root.Right = NewNode(20)
	root.Right.Left = NewNode(15)
	root.Right.Right = NewNode(7)
	res := GivenLevelOrder(root, 2)
	fmt.Println(res)
	res = GivenLevelOrder(root, 1)
	fmt.Println(res)
}
func TestGivenLevelOrderFromRight(t *testing.T) {
	root := NewNode(3)
	root.Left = NewNode(9)
	root.Left.Left = NewNode(8)
	root.Left.Right = NewNode(10)
	root.Right = NewNode(20)
	root.Right.Left = NewNode(15)
	root.Right.Right = NewNode(7)
	res := GivenLevelOrderFromRight(root, 2)
	fmt.Println(res)
}
func TestLevelOrder(t *testing.T) {
	root := NewNode(3)
	root.Left = NewNode(9)
	root.Right = NewNode(20)
	root.Right.Left = NewNode(15)
	root.Right.Right = NewNode(7)
	res := LevelOrder(root)
	fmt.Println(res)
}
func TestLevelOrderBottom(t *testing.T) {
	root := NewNode(3)
	root.Left = NewNode(9)
	root.Right = NewNode(20)
	root.Right.Left = NewNode(15)
	root.Right.Right = NewNode(7)
	res := LevelOrderBottom(root)
	fmt.Println(res)
}
