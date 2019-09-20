package main

import "fmt"

func main() {
	// the empty interface , so i can any type , default is "hello"
	var i interface{} = "hello"
	fmt.Println("i am string: ", i)
	// Interface  Type assertions
	is, ok := i.(string)
	fmt.Println("i am string after type assertions:", is, ok)

	i = 1
	it, _ := i.(int)
	fmt.Println("i am int after type assertions:", it)

	i = true
	ib, _ := i.(int)
	fmt.Println("i am int after type assertions:", ib)
}
