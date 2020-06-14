package singlylinkedlist

import (
	"fmt"
	"testing"
)

func TestList_SingleLinkedList(t *testing.T) {
	singleLinkedList := new(List)
	singleLinkedList.Append(1)
	singleLinkedList.Append(2)
	singleLinkedList.Append(3)
	singleLinkedList.Append(4)
	singleLinkedList.Prepend(0)
	fmt.Println(singleLinkedList.head)
	fmt.Println(singleLinkedList.head.next)
	fmt.Println(singleLinkedList.head.next.next)
	fmt.Println(singleLinkedList.head.next.next.next)
	fmt.Println(singleLinkedList.head.next.next.next.next)
	singleLinkedList.Reverse()
	fmt.Println(singleLinkedList.head)
	fmt.Println(singleLinkedList.head.next)
	fmt.Println(singleLinkedList.head.next.next)
	fmt.Println(singleLinkedList.head.next.next.next)
	fmt.Println(singleLinkedList.head.next.next.next.next)
}
