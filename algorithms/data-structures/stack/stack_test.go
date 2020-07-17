package stack

import (
	"fmt"
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
