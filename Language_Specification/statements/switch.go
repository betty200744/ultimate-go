package main

import "fmt"

func Do(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Println("i am string")
	case int:
		fmt.Println("i am int")
	case bool:
		fmt.Println("i am bool")
	default:
		fmt.Println("i don't know")
	}
}

func main() {
	// no default is bed code

	// value switch
	i := 3
	switch i {
	case 1, 2:
		fmt.Println("this is 1 or 2")
	case 3:
		fmt.Println("this is 3")
	default:
		fmt.Println("this is defalut")
		break
	}

	switch {
	case i < 1:
		fmt.Println(" i < 1")
	default:
		fmt.Println("i < 1")
	}

	// value switch
	Do("string")
	Do(123)
	Do(true)
}
