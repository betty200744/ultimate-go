package __addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(value int) *ListNode {
	return &ListNode{
		Val:  value,
		Next: nil,
	}
}
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &ListNode{}
	cur := l
	curry := 0
	for l1 != nil || l2 != nil || curry > 0 {
		sum := curry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		value := sum % 10
		curry = sum / 10
		cur.Next = &ListNode{
			Val: value,
		}
		cur = cur.Next
	}
	return l.Next
}
