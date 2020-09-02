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
func PreOrderIterative(root *TreeNode) []int {
	q := make([]*TreeNode, 0)
	q = append(q, root)
	n := q[0]
	for n.Left != nil || n.Right != nil {
		if n.Left != nil {
			q = append(q, n.Left)
		}
		if n.Right != nil {
			q = append(q, n.Right)
		}
		n = q[len(q)-1]
	}
	arr := make([]int, len(q))
	for i, i2 := range q {
		arr[i] = i2.Val
	}
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
