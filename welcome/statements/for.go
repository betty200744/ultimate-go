package main

import "fmt"

func main() {
	// 3 type of for
	i := 1
	for i < 3 {
		fmt.Println(i)
		i++
	}

	for i := 1; i < 3; i++ {
		fmt.Println(i)
	}
}
