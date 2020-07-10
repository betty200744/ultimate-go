package singlylinkedlist

import "fmt"

// A list where each element points to the next element in the list.

// 顺序是根据每个对象定的
/*
增，删，改，查
AddFirst or Prepend(value interface{}), O(1)
Append(value interface{}), O(n) , 后面加,
AddLast(value interface{}),  O(1), 后面加, 添加
RemoveFirst(index int) interface{}, O(1)
RemoveLast(index int) interface{}, O(1)
reverse()
Find(node *Node) (int, error), O(n)
Peek(),  O(n)
Contains(values ...interface{}) bool,  O(n)
*/
type Node struct {
	value interface{}
	next  *Node
}
type LinkedList struct {
	head *Node
	tail *Node // 用于addLast
	size int
}

func NewNode(value interface{}) *Node {
	return &Node{value: value, next: nil}
}
func (l *LinkedList) Size() int {
	return l.size
}
func (l *LinkedList) Prepend(value interface{}) {
	if l.head == nil {
		newNode := NewNode(value)
		l.head = newNode
		l.tail = newNode
	} else {
		newNode := NewNode(value)
		newNode.next = l.head
		l.head = newNode
	}
	l.size += 1
}
func (l *LinkedList) Append(value interface{}) {
	newNode := NewNode(value)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		tmp := l.head
		// 因为没有tail， 此时得loop n times查找next是否为最后一个
		for tmp.next != nil {
			tmp = tmp.next
		}
		tmp.next = newNode
		l.tail = newNode
	}
	l.size += 1
}
func (l *LinkedList) AddLast(value interface{}) {
	if l.head == nil {
		l.head = NewNode(value)
		l.tail = l.head
	} else {
		newNode := NewNode(value)
		l.tail.next = newNode
		l.tail = newNode
	}
	l.size += 1
}
func (l *LinkedList) RemoveFirst() interface{} {
	// empty
	if l.head == nil {
		return nil
	}
	tmp := l.head.value
	// single element
	if l.head.next == nil {
		l.head = nil
		l.tail = nil
	} else { // multi element
		l.head = l.head.next
	}
	l.size -= 1
	return tmp
}
func (l *LinkedList) RemoveLast() interface{} {
	if l.head == nil {
		return nil
	}
	if l.head.next == nil {
		return l.RemoveFirst()
	}
	// 找到倒数第二个
	current := l.head
	for ; current.next.next != nil; current = current.next {
	}
	// 取出倒数第一个的value
	tmp := current.next.value
	// 把倒数第二个变成最后一个
	current.next = nil
	l.tail = current
	return tmp
}
func (l *LinkedList) Reverse() {
	var prev, next *Node
	cur := l.head
	for cur != nil {
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	l.head = prev
}
func (l *LinkedList) peekFirst() interface{} {
	if l.head == nil {
		return nil
	}
	return l.head.value
}
func (l *LinkedList) peekLast() interface{} {
	if l.tail == nil {
		return nil
	}
	return l.tail.value
}
func (l *LinkedList) peekLast2() interface{} {
	if l.tail == nil {
		return nil
	}
	tmp := l.head
	for ; tmp.next != nil; tmp = tmp.next {
	}
	return tmp.value
}
func (l *LinkedList) Print() {
	for current := l.head; current != nil; current = current.next {
		fmt.Println(current)
	}
}
