package main

import "fmt"

const co string = "co"

func main() {
	// character , string, int, float, boolean , cannot assign
	const co1 string = "co1"
	const co2 int = 2
	const co3 = true
	fmt.Println(co, co1, co2, co3)
}
