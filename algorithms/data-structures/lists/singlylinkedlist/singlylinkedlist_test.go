package singlylinkedlist

import (
	"fmt"
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestList_SingleLinkedList(t *testing.T) {
	singleLinkedList := new(LinkedList)
	singleLinkedList.Prepend(2)
	singleLinkedList.Prepend(1)
	singleLinkedList.Prepend(0)
	singleLinkedList.Print()
	fmt.Println("this is first: ", singleLinkedList.peekFirst())
	fmt.Println("this is last: ", singleLinkedList.peekLast())
	fmt.Println("this is last2: ", singleLinkedList.peekLast2())
	fmt.Println(singleLinkedList.RemoveLast())
	fmt.Println(singleLinkedList.RemoveFirst())
}
func TestLinkedList_Prepend(t *testing.T) {
	singleLinkedList := new(LinkedList)
	singleLinkedList.Prepend(2)
	singleLinkedList.Prepend(1)
	singleLinkedList.Prepend(0)
	singleLinkedList.Print()
	assert.Equal(t, 0, singleLinkedList.peekFirst())
	assert.Equal(t, 2, singleLinkedList.peekLast())
}

func TestLinkedList_Append(t *testing.T) {
	singleLinkedList2 := new(LinkedList)
	singleLinkedList2.Append(0)
	singleLinkedList2.Append(1)
	singleLinkedList2.Append(2)
	singleLinkedList2.Print()
	assert.Equal(t, 0, singleLinkedList2.peekFirst())
	assert.Equal(t, 2, singleLinkedList2.peekLast())
}
func TestLinkedList_RemoveFirst(t *testing.T) {
	singleLinkedList := new(LinkedList)
	singleLinkedList.Prepend(2)
	singleLinkedList.Prepend(1)
	singleLinkedList.Prepend(0)
	singleLinkedList.Print()
	assert.Equal(t, 0, singleLinkedList.RemoveFirst())
}
func TestLinkedList_RemoveLast(t *testing.T) {
	singleLinkedList := new(LinkedList)
	singleLinkedList.Prepend(2)
	singleLinkedList.Prepend(1)
	singleLinkedList.Prepend(0)
	singleLinkedList.Print()
	assert.Equal(t, 2, singleLinkedList.RemoveLast())
}
func TestLinkedList_Reverse(t *testing.T) {
	singleLinkedList3 := new(LinkedList)
	singleLinkedList3.AddLast(0)
	singleLinkedList3.AddLast(1)
	singleLinkedList3.AddLast(2)
	singleLinkedList3.Print()
	singleLinkedList3.Reverse()
	singleLinkedList3.Print()
}
