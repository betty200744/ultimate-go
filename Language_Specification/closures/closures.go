package main

import "fmt"

func closures() func() int {
	a := 1
	return func() int {
		a++
		return a
	}
}

func main() {
	nexta := closures()
	fmt.Println(nexta())
	fmt.Println(nexta())
	fmt.Println(nexta())
	fmt.Println(nexta())
}
