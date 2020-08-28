package _7_mirrorTree

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

func (t *Node) Insert(data int) *Node {
	if data > t.Data {
		if t.Right != nil {
			t.Right.Insert(data)
		} else {
			t.Right = NewNode(data)
		}
	} else {
		if t.Left != nil {
			t.Left.Insert(data)
		} else {
			t.Left = NewNode(data)
		}

	}
	return t
}

// Breadth-first search
func BreadthFirst(root *Node) {
	var q []*Node
	q = append(q, root)
	for len(q) > 0 {
		n := q[0]
		fmt.Print(n.Data, " ")
		q = q[1:]
		if n.Left != nil {
			q = append(q, n.Left)
		}
		if n.Right != nil {
			q = append(q, n.Right)
		}
	}
}

func CountNode(node *Node) int {
	if node == nil {
		return 0
	}
	return CountNode(node.Left) + CountNode(node.Right) + 1
}
func MirrorTree(root *Node) {
	//if CountNode(root)%2 != 0 {
	//	panic("error")
	//}
	if root.Left == nil && root.Right == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	MirrorTree(root.Left)
	MirrorTree(root.Right)
	return
}
