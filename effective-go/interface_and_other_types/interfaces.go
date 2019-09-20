package main

import (
	"fmt"
	"sort"
)

// interface, special the behavior of object
// interface, a type can implement multiple interface
// interface, 名字一般参考对应的方法， 如write method， interface 就为io.Writer
// interface, can implement multiple interfaces
type Sequence []int

// Methods required by sort.Interface.
func (s *Sequence) Len() int {
	return len(*s)
}
func (s *Sequence) Less(i, j int) bool {
	return (*s)[i] < (*s)[j]
}
func (s *Sequence) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

// Copy returns a copy of the Sequence.
func (s *Sequence) Copy() *Sequence {
	copyy := make(Sequence, 0, len(*s))
	ss := append(copyy, *s...)
	return &ss
}

// Method for printing - sorts the elements before printing.
func (s *Sequence) String() string {
	fmt.Println(sort.IsSorted(s))
	s = s.Copy() // Make a copy; don't overwrite argument.
	sort.Sort(s)
	str := "["
	for i, elem := range *s { // Loop is O(N²); will fix that in next example.
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]"
}

func main() {
	ss := &Sequence{4, 2, 1, 3}
	fmt.Println(ss)
	/*	fmt.Println(ss.Len())
		fmt.Println(ss.Less(1, 2))
		ss.Swap(0, 1)
		fmt.Println(ss)
		fmt.Println(ss.Copy())*/
}
