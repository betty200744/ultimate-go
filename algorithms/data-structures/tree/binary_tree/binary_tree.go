package binary_tree

import "fmt"

/*
 MIT: https://ocw.mit.edu/courses/electrical-engineering-and-computer-science/6-006-introduction-to-algorithms-fall-2011/lecture-videos/MIT6_006F11_lec05.pdf
 Binary Tree: https://en.wikipedia.org/wiki/Binary_tree
 Tree traversal : https://en.wikipedia.org/wiki/Tree_traversal

Node Properties:
Data
Left Node
Right Node

Operations:
Insert
Delete
Find or Contains or Search
FindMin
Traversal, Depth-first search, Pre-order(NLR)
Traversal, Depth-first search, In-order(LNR)
Traversal, Depth-first Post-order (LRN)
Traversal, Breadth-first search
*/

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

func Visit(node *Node) {
	fmt.Print(node.Data, ",")
}

// NLR
func PreOrder(node *Node) {
	if node == nil {
		return
	}
	Visit(node)
	PreOrder(node.Left)
	PreOrder(node.Right)
}

// LNR
func InOrder(node *Node) {
	if node == nil {
		return
	}
	InOrder(node.Left)
	Visit(node)
	InOrder(node.Right)
}

// LRN
func PostOrder(node *Node) {
	if node == nil {
		return
	}
	PostOrder(node.Left)
	PostOrder(node.Right)
	Visit(node)
}

// Breadth-first search
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
