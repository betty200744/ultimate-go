package _26_invertTree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(value int) *TreeNode {
	return &TreeNode{
		Val:   value,
		Left:  nil,
		Right: nil,
	}
}

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	InvertTree(root.Right)
	InvertTree(root.Left)
	return root
}
func BreadthFirst(root *TreeNode) {
	var q []*TreeNode // queue
	q = append(q, root)
	for len(q) != 0 {
		n := q[0]
		q = q[1:]
		fmt.Print(n.Val, " ")
		if n.Left != nil {
			q = append(q, n.Left)
		}
		if n.Right != nil {
			q = append(q, n.Right)
		}
	}
}
