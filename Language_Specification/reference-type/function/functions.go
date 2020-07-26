package main

import "fmt"

/*

function
func()
func(x int) int
func(a, _ int, z float32) bool
func(a, b int, z float32) (bool)
func(prefix string, values ...int)
func(a, b int, z float64, opt ...interface{}) (success bool)
func(int, int, float64) (float64, *[]int)
func(int, int, float64) (*[]int， error)
func(n int) func(p *T)
*/

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
