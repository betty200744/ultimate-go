package queue

/*
定长
first-in , first-out
           +---+---+---+
enqueue -> | 3 | 2 | 1 | -> dequeue
           +---+---+---+
*/
/*
	Enqueue(value interface{})
	Dequeue() interface{}
*/
type Queue struct {
	queue []interface{}
	len   int
}

func New() *Queue {
	queue := &Queue{
		queue: []interface{}{},
		len:   0,
	}
	return queue
}

func (queue *Queue) Enqueue(value interface{}) {
	queue.queue = append(queue.queue, value)
	queue.len++
}
func (queue *Queue) Dequeue() interface{} {
	el := queue.queue[0]
	queue.queue = queue.queue[1:]
	queue.len--
	return el
}
