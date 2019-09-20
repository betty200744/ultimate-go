package main

import (
	"flag"
	"fmt"
)

var (
	wordP string
	numbP int
)

func main() {

	flag.StringVar(&wordP, "word", "", "input a string")
	flag.IntVar(&numbP, "num", 0, "input a number")

	flag.Parse()
	fmt.Println("word: ", wordP)
	fmt.Println("numb: ", numbP)

	a := []string{}
	fmt.Println(len(a))
}
