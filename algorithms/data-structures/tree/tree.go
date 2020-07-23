package tree

import "fmt"

type Node struct {
	Left  *Node
	Right *Node
	Data  int
}

func New(data int) *Node {
	return &Node{
		Left:  nil,
		Right: nil,
		Data:  data,
	}
}

func (t *Node) Add(data int) *Node {
	if data > t.Data {
		if t.Right != nil {
			t.Right.Add(data)
		} else {
			t.Right = New(data)
		}
	} else {
		if t.Left != nil {
			t.Left.Add(data)
		} else {
			t.Left = New(data)
		}

	}
	return t
}
func (t *Node) Contains(data int) bool {
	if t.Data == data {
		return true
	}
	if data < t.Data {
		if t.Left != nil {
			return t.Left.Contains(data)
		}
	} else {
		if t.Right != nil {
			return t.Right.Contains(data)
		}
	}
	return false
}

func (t *Node) Print() {
	fmt.Print(t.Data, ", ")
	if t.Right != nil {
		t.Right.Print()
	}
	if t.Left != nil {
		t.Left.Print()
	}
}
