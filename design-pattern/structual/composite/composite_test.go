package composite

import "testing"

func TestNewNode(t *testing.T) {
	root := NewNode(0)
	c1 := NewNode(1)
	c2 := NewNode(2)
	root.AddChild(c1)
	root.AddChild(c2)
	root.Print()
	for _, i2 := range root.children {
		i2.Print()
	}
}
