package main

import "fmt"

func main() {
	// string , integer, float, boolean
	fmt.Println("a" + "b")
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1.3 + 1, 3 =", 1.3+1.3)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!!true)
}
