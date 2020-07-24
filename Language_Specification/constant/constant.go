package main

import "fmt"

const (
	A1 = iota // 如同数据库里面的serial, 自增
	A2 = iota
	A3 = iota
)

func main() {
	fmt.Printf("A1 :  %T, [%v]", A1, A1)
}
