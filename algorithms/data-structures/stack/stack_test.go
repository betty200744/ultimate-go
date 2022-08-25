package stack

import (
	"fmt"
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func Test_Stacks(t *testing.T) {
	s := New()
	s.Push("a")
	s.Push("b")
	s.Push("c")
	s.Push("d")
	fmt.Println(s)
	fmt.Println(s.Peek())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.IsEmpty())
}
func TestStack_Max(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		s := New()
		s.Push(1)
		s.Push(2)
		s.Push(3)
		s.Push(4)
		assert.Equal(t, 4, s.Max())
	})
	t.Run("string", func(t *testing.T) {
		s := New()
		s.Push("a")
		s.Push("b")
		s.Push("c")
		s.Push("d")
		assert.Equal(t, "d", s.Max())
	})
}
