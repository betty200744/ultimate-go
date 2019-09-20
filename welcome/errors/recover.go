package errors

import "fmt"

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(i)
	}
	g(i + 1)
}

func ff() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	g(0)
	fmt.Println("Returned normally from g.")
}

func main() {
	ff()
	fmt.Println("Returned normally from f.")
}
