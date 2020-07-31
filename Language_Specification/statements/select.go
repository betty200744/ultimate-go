package main

import "fmt"

func main() {
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

		var a string
		select {
		case <-ch1: // recvExpr
		case a = <-ch1: // expressionList
		case b := <-ch1:
			{ //IdentifierList
				fmt.Println(b)
			}
		}
	}
}
