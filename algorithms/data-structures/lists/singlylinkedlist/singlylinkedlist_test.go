package singlylinkedlist

import (
	"fmt"
	"testing"
)

func TestList_SingleLinkedList(t *testing.T) {
	singleLinkedList := new(LinkedList)
	singleLinkedList.Prepend(2)
	singleLinkedList.Prepend(1)
	singleLinkedList.Prepend(0)
	fmt.Println(singleLinkedList.head, singleLinkedList.head.next, singleLinkedList.head.next.next, singleLinkedList.size)
	fmt.Println(singleLinkedList.RemoveLast())
	fmt.Println(singleLinkedList.RemoveFirst())
	singleLinkedList2 := new(LinkedList)
	singleLinkedList2.Append(0)
	singleLinkedList2.Append(1)
	singleLinkedList2.Append(2)
	fmt.Println(singleLinkedList2.head, singleLinkedList2.head.next, singleLinkedList2.head.next.next, singleLinkedList2.size)
	singleLinkedList3 := new(LinkedList)
	singleLinkedList3.AddLast(0)
	singleLinkedList3.AddLast(1)
	singleLinkedList3.AddLast(2)
	fmt.Println(singleLinkedList3.head, singleLinkedList3.head.next, singleLinkedList3.head.next.next, singleLinkedList3.size)
	singleLinkedList3.Reverse()
	fmt.Println(singleLinkedList3.head, singleLinkedList3.head.next, singleLinkedList3.head.next.next, singleLinkedList3.size)
}
