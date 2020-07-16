package queue

/*
定长
first-in , first-out
           +---+---+---+
enqueue -> | 3 | 2 | 1 | -> dequeue
           +---+---+---+
*/
/*
	IsEmpty() bool
	Peek() (interface{}, error)
	Enqueue(value interface{})
	Dequeue() interface{}
*/
/*
现实场景如： 消息队列，kafka, rabbitmq
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
func (q *Queue) IsEmpty() bool {
	return false
}
func (q *Queue) Peek() (interface{}, bool) {
	return nil, false
}
func (q *Queue) Enqueue(value interface{}) {
	q.queue = append(q.queue, value)
	q.len++
}
func (q *Queue) Dequeue() interface{} {
	el := q.queue[0]
	q.queue = q.queue[1:]
	q.len--
	return el
}
