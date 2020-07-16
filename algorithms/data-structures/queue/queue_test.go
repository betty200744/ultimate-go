package queue

import (
	"fmt"
	"testing"
)

func Test_Queue(t *testing.T) {
	queue := New()
	queue.Enqueue("a")
	queue.Enqueue("b")
	fmt.Println(queue.queue)
	queue.Dequeue()
	fmt.Println(queue.queue)
	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.Peek())
}
