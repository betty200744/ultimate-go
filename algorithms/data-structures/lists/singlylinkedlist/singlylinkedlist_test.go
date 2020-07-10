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
	singleLinkedList.Print()
	fmt.Println("this is first: ", singleLinkedList.peekFirst())
	fmt.Println("this is last: ", singleLinkedList.peekLast())
	fmt.Println("this is last2: ", singleLinkedList.peekLast2())
	fmt.Println(singleLinkedList.RemoveLast())
	fmt.Println(singleLinkedList.RemoveFirst())
	singleLinkedList2 := new(LinkedList)
	singleLinkedList2.Append(0)
	singleLinkedList2.Append(1)
	singleLinkedList2.Append(2)
	singleLinkedList2.Print()
	fmt.Println("this is first: ", singleLinkedList2.peekFirst())
	fmt.Println("this is last: ", singleLinkedList2.peekLast())
	fmt.Println("this is last2: ", singleLinkedList2.peekLast2())
	singleLinkedList3 := new(LinkedList)
	singleLinkedList3.AddLast(0)
	singleLinkedList3.AddLast(1)
	singleLinkedList3.AddLast(2)
	fmt.Println("this is first: ", singleLinkedList3.peekFirst())
	fmt.Println("this is last: ", singleLinkedList3.peekLast())
	fmt.Println("this is last2: ", singleLinkedList3.peekLast2())
	singleLinkedList3.Print()
	singleLinkedList3.Reverse()
	singleLinkedList3.Print()
}
