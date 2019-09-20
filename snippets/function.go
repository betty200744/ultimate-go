package main

import "fmt"

// no parameter function
func aoo() {
	fmt.Println("no parameter function")
}

// has parameter function
func boo(x int) {
	fmt.Println("has parameter function, arg x :", x)
}

// has return function
func coo(x int) int {
	fmt.Println("has return function, return :", x)
	return x
}

// closures function
func doo() func() int {
	i := 2
	j := func() int { return i }
	return j
}

func main() {
	aoo()
	boo(1)
	fmt.Println(coo(1))
	fmt.Println(doo()())
}
