package container

import (
	"container/list"
	"fmt"
	"testing"
)

// doubly linked list

func Test_List(t *testing.T) {
	l := list.New()
	l1 := l.PushBack(1)
	l2 := l.PushBack(2)
	l0 := l.PushFront(0)
	l.InsertAfter(3, l2)
	fmt.Println("this is list front:", l.Front())
	fmt.Println("this is list back:", l.Back())
	fmt.Println("this is l0:", l0.Value)
	fmt.Println("this is l1:", l1.Value)
	fmt.Println("this is len:", l.Len())
}
