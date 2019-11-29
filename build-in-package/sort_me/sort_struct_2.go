package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name  string
	Grade int
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].Grade < s[j].Grade
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	students := Students{
		{"Denis", 12},
		{"Ascolin", 50},
		{"IIsus", 99},
		{"Mario", 2},
		{"Gogaie", 37},
		{"Dentistul", 76},
	}

	fmt.Println(students)
	sort.Sort(students)

	fmt.Println(students)
	fmt.Println(sort.IsSorted(students))

	sort.Sort(sort.Reverse(students))
	fmt.Println(students)

}
