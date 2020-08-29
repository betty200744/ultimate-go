package HasPathSum

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

// recurse
func HasPathSum(root *Node, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Data == sum
	}
	return HasPathSum(root.Left, sum-root.Data) || HasPathSum(root.Right, sum-root.Data)
}
