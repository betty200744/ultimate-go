package stacks

// https://en.wikipedia.org/wiki/Stack_%28abstract_data_type%29#Array
// stacks 如同叠盘子， 先进后出

/*
	go 语言中就是Array来实现, Slice来实现
	last-in, first-out
	Push(value interface{})
	Pop() ()
*/
/*
现实场景， 代码的压栈
*/
type Stack struct {
	stack []interface{}
	len   int
}

func New() *Stack {
	s := &Stack{
		stack: make([]interface{}, 0),
		len:   0,
	}
	return s
}
func (stack *Stack) Push(value interface{}) {
	stack.stack = append([]interface{}{value}, stack.stack...)
	stack.len++
}
func (stack *Stack) Pop() interface{} {
	el := stack.stack[0]
	stack.stack = stack.stack[1:]
	stack.len--
	return el
}
