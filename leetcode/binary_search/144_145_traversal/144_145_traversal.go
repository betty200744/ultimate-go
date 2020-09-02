package _44_145_traversal

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

// NLR
func PreOrder(root *TreeNode) []int {
	arr := make([]int, 0)
	if root == nil {
		return arr
	}
	arr = append(arr, root.Val)
	arr = append(arr, PreOrder(root.Left)...)
	arr = append(arr, PreOrder(root.Right)...)
	return arr
}

// LNR
func InOrder(root *TreeNode) []int {
	arr := make([]int, 0)
	if root == nil {
		return arr
	}
	arr = append(arr, InOrder(root.Left)...)
	arr = append(arr, root.Val)
	arr = append(arr, InOrder(root.Right)...)
	return arr
}

// LRN
func PostOrder(root *TreeNode) []int {
	arr := make([]int, 0)
	if root == nil {
		return arr
	}
	arr = append(arr, PostOrder(root.Left)...)
	arr = append(arr, PostOrder(root.Right)...)
	arr = append(arr, root.Val)
	return arr
}
