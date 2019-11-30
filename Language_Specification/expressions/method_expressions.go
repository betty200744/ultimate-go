package main

import "fmt"

// struct has method set
type MM struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (m *MM) GetNameById(id int64) *MM {
	return m
}

func (m *MM) GetNameByName(name string) *MM {
	return m
}

func main() {
	m1 := &MM{Id: 1, Name: "betty"}
	fmt.Println(m1.GetNameById(1)) // method expressions
	fmt.Println(m1.GetNameByName("betty"))
}
