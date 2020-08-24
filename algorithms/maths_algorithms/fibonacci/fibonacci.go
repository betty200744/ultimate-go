package main

import "fmt"

func fibonacci() func() int {
	num1 := 0
	num2 := 1
	return func() int {
		num := num1
		num1, num2 = num2, num+num2
		return num
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
