package main

import (
	"fmt"
	"os"
)

func main() {
	// 3 type if else , this is no ternary in go
	if 4%2 == 0 {
		fmt.Println("this is if")
	}

	if 4%2 == 0 {
		fmt.Println("this is if else")
	} else {
		fmt.Println("this is if else")
	}

	if 4%2 == 0 {
		fmt.Println("this is if else if")
	} else if 4%3 == 0 {
		fmt.Println("this is if else if")
	} else {
		fmt.Println("this is if else if")
	}

	// bad code
	groupsBad, errBad := os.Getgroups()
	if errBad == nil {
		fmt.Println(groupsBad)
	}

	// good code
	if groups, err := os.Getgroups(); err != nil {
		for _, g := range groups {
			fmt.Println(g)
		}
	}

}
