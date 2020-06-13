package main

import "fmt"

type List struct {
	elements []interface{}
	size     int
}
type Iterator struct {
	list  *List
	index int
}

func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}
func (list *List) Iterator() Iterator {
	return Iterator{list: list, index: -1}
}
func (iterator *Iterator) Next() bool {
	if iterator.index < iterator.list.size {
		iterator.index++
	}
	return iterator.index >= 0 && iterator.index < iterator.list.size
}
func (iterator *Iterator) Index() int {
	return iterator.index
}
func (iterator *Iterator) Value() interface{} {
	return iterator.list.elements[iterator.index]
}
func (list *List) Add(values ...interface{}) {
	n := len(values)
	currentCapacity := cap(list.elements)
	if list.size+n >= currentCapacity {
		newCapacity := 2 * (currentCapacity + n)
		newElements := make([]interface{}, newCapacity, newCapacity)
		copy(newElements, list.elements)
		list.elements = newElements
	}
	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}
}
func (list *List) Size() int {
	return list.size
}
func (list *List) Empty() bool {
	if list.size == 0 {
		return true
	}
	return false
}
func (list *List) Values() []interface{} {
	return list.elements
}
func (list *List) Contains(values ...interface{}) bool {
	found := false
	for _, i2 := range values {
		for _, i3 := range list.elements {
			if i2 == i3 {
				found = true
				break
			}
		}
	}
	return found
}
func (list *List) Set(index int, value interface{}) {
	if index >= 0 && index < list.size {
		list.elements[index] = value
	} else {
		list.Add(value)
	}
}
func (list *List) Each(f func(index int, value interface{})) {
	iterator := list.Iterator()
	for iterator.Next() {
		f(iterator.index, iterator.Value())
	}
}
func (list *List) Find(f func(index int, value interface{}) bool) (int, interface{}) {
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.index, iterator.Value()) {
			return iterator.Index(), iterator.Value()
		}
	}
	return -1, nil
}

func main() {
	list := New(0, 1, 2)
	list.Add(1)
	fmt.Println(list.elements)
	fmt.Println(list.size)
	fmt.Println(list.Empty())
	fmt.Println(list.Values())
	fmt.Println(list.Contains(1))
	fmt.Println(list.Contains(5))
	list.Set(3, 3)
	list.Set(6, 4)
	fmt.Println(list.Values())
	list.Each(func(index int, value interface{}) {
		fmt.Println(index, value)
	})
	fmt.Println(list.Find(func(index int, value interface{}) bool {
		return true
	}))
}
