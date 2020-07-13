package singlylinkedlist

import "fmt"

// A list where each element points to the Next element in the list.

// 顺序是根据每个对象定的
/*
增，删，改，查
AddFirst or Prepend(Value interface{}), O(1)
Append(Value interface{}), O(n) , 后面加,
AddLast(Value interface{}),  O(1), 后面加, 添加
RemoveFirst(index int) interface{}, O(1)
RemoveLast(index int) interface{}, O(1)
reverse()
Find(node *Node) (int, error), O(n)
Peek(),  O(n)
Contains(values ...interface{}) bool,  O(n)
*/
type Node struct {
	Value interface{}
	Next  *Node
}
type LinkedList struct {
	Head *Node
	tail *Node // 用于addLast
	size int
}

func NewNode(value interface{}) *Node {
	return &Node{Value: value, Next: nil}
}
func (l *LinkedList) Size() int {
	return l.size
}
func (l *LinkedList) Prepend(value interface{}) {
	if l.Head == nil {
		newNode := NewNode(value)
		l.Head = newNode
		l.tail = newNode
	} else {
		newNode := NewNode(value)
		newNode.Next = l.Head
		l.Head = newNode
	}
	l.size += 1
}
func (l *LinkedList) Append(value interface{}) {
	newNode := NewNode(value)
	if l.Head == nil {
		l.Head = newNode
		l.tail = newNode
	} else {
		tmp := l.Head
		// 因为没有tail， 此时得loop n times查找next是否为最后一个
		for tmp.Next != nil {
			tmp = tmp.Next
		}
		tmp.Next = newNode
		l.tail = newNode
	}
	l.size += 1
}
func (l *LinkedList) AddLast(value interface{}) {
	if l.Head == nil {
		l.Head = NewNode(value)
		l.tail = l.Head
	} else {
		newNode := NewNode(value)
		l.tail.Next = newNode
		l.tail = newNode
	}
	l.size += 1
}
func (l *LinkedList) RemoveFirst() interface{} {
	// empty
	if l.Head == nil {
		return nil
	}
	tmp := l.Head.Value
	// single element
	if l.Head.Next == nil {
		l.Head = nil
		l.tail = nil
	} else { // multi element
		l.Head = l.Head.Next
	}
	l.size -= 1
	return tmp
}
func (l *LinkedList) RemoveLast() interface{} {
	if l.Head == nil {
		return nil
	}
	if l.Head.Next == nil {
		return l.RemoveFirst()
	}
	// 找到倒数第二个
	current := l.Head
	for ; current.Next.Next != nil; current = current.Next {
	}
	// 取出倒数第一个的value
	tmp := current.Next.Value
	// 把倒数第二个变成最后一个
	current.Next = nil
	l.tail = current
	return tmp
}
func (l *LinkedList) Reverse() {
	var prev, next *Node
	cur := l.Head
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	l.Head = prev
}
func (l *LinkedList) peekFirst() interface{} {
	if l.Head == nil {
		return nil
	}
	return l.Head.Value
}
func (l *LinkedList) peekLast() interface{} {
	if l.tail == nil {
		return nil
	}
	return l.tail.Value
}
func (l *LinkedList) peekLast2() interface{} {
	if l.tail == nil {
		return nil
	}
	tmp := l.Head
	for ; tmp.Next != nil; tmp = tmp.Next {
	}
	return tmp.Value
}
func (l *LinkedList) Print() {
	for current := l.Head; current != nil; current = current.Next {
		fmt.Println(current)
	}
}
