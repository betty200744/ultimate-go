package _43_diameterOfBinaryTree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{
		Val:   value,
		Left:  nil,
		Right: nil,
	}
}
func Height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lheight := 1
	rheight := 1
	if root.Left != nil {
		lheight = Height(root.Left) + 1
	}
	if root.Right != nil {
		rheight = Height(root.Right) + 1
	}
	return int(math.Max(float64(lheight), float64(rheight)))
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func DiameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lheight := Height(root.Left)
	rheight := Height(root.Right)
	rootDiameter := lheight + rheight
	lDiameter := DiameterOfBinaryTree(root.Left)
	rDiameter := DiameterOfBinaryTree(root.Right)
	return max(rootDiameter, max(lDiameter, rDiameter))
}
