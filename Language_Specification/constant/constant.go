package main

import "fmt"

const (
	A1 = iota // 如同数据库里面的serial, 自增, iota默认Increment by 1
	A2
	A3
)
const (
	B1 = iota + 2 // start by iot + 2,即2, Increment by iot, 即1
	B2
	B3
)
const (
	C1 = iota * 2 // start iot*2, 即0, Increment by iot * 2, 即2
	C2
	C3
	C4
)
const (
	D1 = 1 << iota // start by 1 , Increment by iot * iot
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
