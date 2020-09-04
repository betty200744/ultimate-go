package _2_getIntersectionNode

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	intersect := NewListNode(8)
	headA := NewListNode(4)
	headA.Next = NewListNode(1)
	headA.Next.Next = intersect
	headB := NewListNode(5)
	headB.Next = NewListNode(3)
	headB.Next.Next = intersect
	res := GetIntersectionNode(headA, headB)
	assert.Equal(t, intersect, res)
}
