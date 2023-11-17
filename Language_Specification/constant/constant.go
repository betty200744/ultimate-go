package main

import "fmt"

const (
	A1 = iota // like mysql serial, auto increment, iota default increment by 1
	A2
	A3
)
const (
	B1 = iota + 2 // start by iot + 2, that is 2, Increment by iot, that is 1
	B2
	B3
)
const (
	C1 = iota * 2 // start iot*2, that 0, Increment by iot * 2, that is 2
	C2
	C3
	C4
)
const (
	D1 = 1 << 2 // start by 1 , left shift 1
	D2
	D3
	D4
)

func main() {
	fmt.Printf("A1 :  %T, [%v] \n", A1, A1)
	fmt.Println(A1, A2, A3)
	fmt.Println(B1, B2, B3)
	fmt.Println(C1, C2, C3, C4)
	fmt.Println(D1, D2, D3, D4)
}
