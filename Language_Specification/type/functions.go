package main

import "fmt"

func returnOne(a, b int) int {
	return a + b
}

func returnTwo(a, b int) (int, int) {
	return a, b
}

func variadic(nums ...interface{}) int {
	var total int
	fmt.Println(nums)
	for _, value := range nums {
		total += value.(int)
	}
	return total
}

func main() {
	// fixed function,  variadic function, one return, two return,
	fmt.Println(returnOne(1, 3))
	fmt.Println(returnTwo(1, 3))

	nums := []interface{}{1, 2, 3}
	fmt.Println(variadic(1, 2, 3))
	fmt.Println(variadic(nums...))
}
