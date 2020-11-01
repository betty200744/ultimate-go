package statements

import (
	"fmt"
	"os"
)

func for_statements() {
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

func if_else_statements() {
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

func select_statements() {
	ch1 := make(chan string, 1)
	ch2 := make(chan int, 1)

	ch1 <- "hello"
	ch2 <- 1
	for i := 1; i < 3; i++ {
		// always no key
		// 只case send Channel or receive Channel
		select {
		case msg := <-ch1:
			fmt.Println("this from ch1" + msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}

		select {
		case <-ch1: // recvExpr
		case b := <-ch1:
			{ //IdentifierList
				fmt.Println(b)
			}
		}
	}
}

func Switch(v interface{}) {
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
