package composite

import "fmt"

type Node struct {
	children []*Node
	data     int
}

func NewNode(data int) *Node {
	return &Node{
		children: nil,
		data:     data,
	}
}

func (n *Node) AddChild(node *Node) {
	n.children = append(n.children, node)
}
func (n *Node) Print() {
	fmt.Println(n.data)
}
