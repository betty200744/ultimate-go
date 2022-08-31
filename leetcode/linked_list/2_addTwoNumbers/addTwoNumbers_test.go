package __addTwoNumbers

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := NewListNode(2)
	l1.Next = NewListNode(4)
	l1.Next.Next = NewListNode(3)
	l2 := NewListNode(5)
	l2.Next = NewListNode(6)
	l2.Next.Next = NewListNode(4)
	res := AddTwoNumbers(l1, l2)
	assert.Equal(t, res.Val, 7)
	assert.Equal(t, res.Next.Val, 0)
	assert.Equal(t, res.Next.Next.Val, 8)
}
