package _2_getIntersectionNode

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
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	for headA != nil {
		tmpB := headB
		for tmpB != nil {
			if headA == tmpB {
				return tmpB
			}
			tmpB = tmpB.Next
		}
		headA = headA.Next
	}
	return nil
}
