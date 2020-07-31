package main

import "fmt"

func main() {
	// for is whale
	i := 1
	for i < 3 {
		fmt.Println(i)
		i++
	}
	// traditional for
	for i := 1; i < 3; i++ {
		fmt.Println(i)
	}
	// for range
	for _, value := range []string{"a", "b"} {
		value = fmt.Sprintf("%s:%s", "prefix", value)
	}
	// for goroutine
	for i := 1; i < 10; i++ {
		fmt.Println("for: ", i)
		go func() {
			fmt.Println("go", i)
		}()
	}
}
