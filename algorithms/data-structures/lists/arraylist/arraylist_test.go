package arraylist

import (
	"fmt"
	"testing"
)

func TestArrayList(t *testing.T) {
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
