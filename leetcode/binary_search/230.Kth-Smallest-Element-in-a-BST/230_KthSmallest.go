package KthSmallest

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

func CountNode(node *Node) int {
	if node == nil {
		return 0
	}
	return CountNode(node.Left) + CountNode(node.Right) + 1
}
func KthSmallest(root *Node, k int) int {
	leftCount := CountNode(root.Left)
	if k == leftCount+1 {
		return root.Data
	}
	if k <= leftCount {
		return KthSmallest(root.Left, k)
	} else {
		return KthSmallest(root.Right, k-leftCount-1)
	}
}
