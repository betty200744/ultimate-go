package main

import "fmt"

func main() {
	// index, array
	// index, pointer to array
	// index, slice
	// index, map
	// index, struct
	// index, string
	// struct, no index

	array := []string{"a", "b"}
	fmt.Println(array[0])

	pointerArray := &array
	fmt.Println((*pointerArray)[0])

	sli := [5]string{"a", "b", "c", "d", "e"}
	sliOdd := sli[0:1:3]
	fmt.Println(sli[0])
	fmt.Println(sliOdd[0])

	mm := make(map[string]string)
	mm["a"] = "a"
	mm["b"] = "b"
	fmt.Println(mm["a"])

	ss := struct {
		A string
		B string
	}{A: "a", B: "b"}
	fmt.Println(ss.A)
}
