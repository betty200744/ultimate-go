package main

import "fmt"

var (
	a  *A
	aa A
)

type A struct {
	Id   string
	Name string
}

func ppointer(pt *[]int) {
	fmt.Println("this is pt:", pt)    // point
	fmt.Println("this is *pt: ", *pt) // value
	*pt = append(*pt, 4)              // change slice p
}

func main() {
	p := []int{1, 2, 3}
	pp := 1
	ppointer(&p)

	fmt.Println("this is p:", p)      //  slice p
	fmt.Println("this is &p:", &p)    // &p , p address, is a point
	fmt.Println("this is &pp: ", &pp) // &pp, pp address, is a point
	fmt.Println(fmt.Sprintf("%s", &A{Id: "id", Name: "name"}))
	fmt.Println(a.Id) // will panic
}
