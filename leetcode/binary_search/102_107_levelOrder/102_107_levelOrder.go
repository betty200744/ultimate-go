package _02_107_levelOrder

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
func Height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lheight := 0
	rheight := 0
	if root.Left != nil {
		lheight = Height(root.Left)
	}
	if root.Right != nil {
		rheight = Height(root.Right)
	}
	if lheight > rheight {
		return lheight + 1
	} else {
		return rheight + 1
	}
}
func LevelOrder(root *TreeNode) [][]int {
	level := Height(root)
	arr := make([][]int, level)
	for i := 0; i < level; i++ {
		arr[i] = GivenLevelOrder(root, i)
	}
	return arr
}
func LevelOrderBottom(root *TreeNode) [][]int {
	level := Height(root)
	arr := make([][]int, level)
	for i := 0; i < level; i++ {
		arr[level-i-1] = GivenLevelOrder(root, i)
	}
	return arr
}
func GivenLevelOrder(root *TreeNode, level int) []int {
	arr := make([]int, 0)
	if root == nil {
		return arr
	}
	if level == 0 {
		return append(arr, root.Val)
	}
	if level > 0 {
		arr = append(arr, GivenLevelOrder(root.Left, level-1)...)
		arr = append(arr, GivenLevelOrder(root.Right, level-1)...)
	}
	return arr
}
