package singlylinkedlist

// A list where each element points to the next element in the list.

// 顺序是根据每个对象定的
/*
Get(index int) interface{}
Set(index int, value interface{})
Add(values ...interface{})
Remove(index int)
Insert(index int, value interface{})
Contains(values ...interface{}) bool
Swap(index1, index2 int)
Prepend(value interface{})
Append(value interface{})
Find(node *Node) (int, error)
Concat(k *List)
reverse()
*/

type Node struct {
	value interface{}
	next  *Node
}
type List struct {
	head *Node
}

func newNode(value interface{}) *Node {
	return &Node{value: value, next: nil}
}

func (list *List) Append(value interface{}) {
	n := newNode(value)
	if list.head == nil {
		list.head = n
		return
	}
	cur := list.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = n
}

// 我在前面， 那么原本的就在我后面
func (list *List) Prepend(value interface{}) {
	n := newNode(value)
	n.next = list.head
	list.head = n
}
func (list *List) Reverse() {
	// 为撒要三个pointer, 因为prev和cur是要换的两个， 另外一个是当前的next 需要存
	var prev, next *Node
	cur := list.head
	for cur != nil {
		// before changing the next of current, store next node
		next = cur.next
		// changing the next of current, 每一个的下一个节点都变成前一个节点
		// 注意此时linked list是断开的， 需要全部都换完才连上
		cur.next = prev
		// move prev and next step, 注意先移prev再移cur，
		prev, cur = cur, next
	}
	// head变成原链表的最后一个
	list.head = prev
}
