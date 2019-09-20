package main

import (
	"fmt"
	"gobyexample/welcome/greet"
	"sort"
)

func main() {
	//  init array, declare array, initial array, initial two Dimensional array, multi-line initial array
	//  set array, get array, len of array , comparison array, append, sort, iteration array for range
	var a [3]int     // declare
	var aa [2][2]int // declare twoD array
	aaa := []string{
		"c",
		"b",
	} //  initial value, longhand var a [n]T = [n]T{V1,V2,...,Vn} , shorthand: a =: [n]T{V1, V2}

	var aaaa []string //declare array
	var aaaaa = new(int)
	aaaaaa := [3]int{1, 2, 3}
	aaaaaaa := [...]string{"aaa", "bbb"}
	aaaaaaaa := [][]string{{"a1", "a2"}, {"b1", "b2"}}

	a[0] = 2     // value assignment
	aa[0][0] = 2 // value assignment
	aaa = append(aaa, "a")
	sort.Strings(aaa)
	copy(aaaa, aaa)
	for e := range aaa {
		fmt.Println(e)
	}

	greett.Greet()

	fmt.Println(a, a[0], aa, aa[0][0], len(a), aaa, aaaa, aaaaa, aaaaaa, aaaaaaa, aaaaaaaa, a == aaaaaa)
}
