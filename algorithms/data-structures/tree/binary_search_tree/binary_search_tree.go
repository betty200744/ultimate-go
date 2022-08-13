package binary_search_tree

import "fmt"

/*
root,
left,
right,

Insert, recursive add
Delete,
Find,
Rotation
Depth

Traversal, Depth-first search, Pre-order(NLR)
Traversal, Depth-first search, In-order(LNR)
Traversal, Depth-first Post-order (LRN)
Traversal, Breadth-first search

inorder successor, is the largest thing smaller then the node
inorder predessor, is the smallest thing large then the node

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
func (t *Node) InOrderSuccessor() *Node {
	cur := t
	if cur.Left != nil {
		cur = cur.Left
	}
	return cur
}

func (t *Node) Delete(data int) *Node {
	// remove a node with two child, swap with inorder successor or predessor then delete the leaf
	if t == nil {
		return nil
	}
	if data < t.Data {
		t.Left = t.Left.Delete(data)
		return t
	}
	if data > t.Data {
		t.Right = t.Right.Delete(data)
		return t
	}
	// remove the leaf, just set parent to null
	if t.Left == nil && t.Right == nil {
		t = nil
		return t
	}
	// remove a node with one child, set parent to the child
	if t.Left == nil {
		t = t.Right
		return t
	}
	if t.Right == nil {
		t = t.Left
		return t
	}
	successor := t.InOrderSuccessor()
	t.Data = successor.Data
	t.Left = t.Left.Delete(data)
	return t
}
func (t *Node) Find(data int) bool {
	if t.Data == data {
		return true
	}
	if data < t.Data {
		if t.Left != nil {
			return t.Left.Find(data)
		}
	} else {
		if t.Right != nil {
			return t.Right.Find(data)
		}
	}
	return false
}
func (t *Node) FindMin() int {
	if t.Left == nil {
		fmt.Println(t.Data)
		return t.Data
	}
	return t.Left.FindMin()
}
func (t *Node) FindMax() int {
	if t.Right == nil {
		fmt.Println(t.Data)
		return t.Data
	}
	return t.Right.FindMax()
}

// Breadth-first search
func BreadthFirst(root *Node) {
	var q []*Node // queue
	var n *Node   // temporary node
	q = append(q, root)
	for len(q) != 0 {
		n, q = q[0], q[1:]
		fmt.Print(n.Data, " ")
		if n.Left != nil {
			q = append(q, n.Left)
		}
		if n.Right != nil {
			q = append(q, n.Right)
		}
	}
}
