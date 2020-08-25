package stack

import "reflect"

// https://en.wikipedia.org/wiki/Stack_%28abstract_data_type%29#Array
// stacks 如同叠盘子， 先进后出

/*
	go 语言中就是Array来实现, Slice来实现
	last-in, first-out
    Motivation， 代码的压栈
	IsEmpty() bool
	Peek() (interface{}, error)
	Push(value interface{})
	Pop() ()
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
func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}
func (s *Stack) Peek() (interface{}, bool) {
	if len(s.stack) == 0 {
		return nil, false
	}
	return s.stack[0], true
}
func (s *Stack) Push(value interface{}) {
	s.stack = append([]interface{}{value}, s.stack...)
	s.len++
}
func (s *Stack) Pop() interface{} {
	el := s.stack[0]
	s.stack = s.stack[1:]
	s.len--
	return el
}
func (s *Stack) Max() interface{} {
	if s.IsEmpty() {
		return 0
	}
	max := s.stack[0]
	for i := 0; i < s.len; i++ {
		switch v := reflect.ValueOf(s.stack[i]); v.Kind() {
		case reflect.String:
			if s.stack[i].(string) > max.(string) {
				max = s.stack[i]
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if s.stack[i].(int) > max.(int) {
				max = s.stack[i]
			}
		default:
			max = 0
		}
	}
	return max
}
