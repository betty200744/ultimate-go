package _99_Binary_Tree_Right_Side_View

import "fmt"

type Node struct {
	Left  *Node
	Right *Node
	Data  int
}

func NewNode(data int) *Node {
	return &Node{
		Left:  nil,
		Right: nil,
		Data:  data,
	}
}
func BreadthFirst(root *Node) {
	var q []*Node // queue
	q = append(q, root)
	for len(q) != 0 {
		n := q[0]
		q = q[1:]
		fmt.Print(n.Data, " ")
		if n.Left != nil {
			q = append(q, n.Left)
		}
		if n.Right != nil {
			q = append(q, n.Right)
		}
	}
}

func RightSideView(root *Node) []int {
	arr := make([]int, 0)
	arr = append(arr, root.Data)
	n := root
	for n.Right != nil {
		arr = append(arr, n.Right.Data)
		n = n.Right
	}
	return arr
}

func LeftSideView(root *Node) []int {
	arr := make([]int, 0)
	arr = append(arr, root.Data)
	n := root
	for n.Left != nil {
		arr = append(arr, n.Left.Data)
		n = n.Left
	}
	return arr
}
