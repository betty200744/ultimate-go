package main

import (
	"fmt"
	"reflect"
)

/*

struct field : class properties
struct method: class behaviors
struct method: class behaviors， 如此一来我们变可以添加相同名字的method，因为receive不同
struct field function: class behaviors
struct type definition 与method definition需要在同一package
*/

type FieldFunction func(s string) string
type StructFieldMethod struct {
	F1 int64
	F2 string
	M1 FieldFunction // 为struct添加method的方式一， 一般不这么做，不符合OOP的概念
}

func M1(n string) string {
	fmt.Println(n)
	return n
}
func (s *StructFieldMethod) ReceiveMethod1(n string) { // 为struct添加method的方式二
	fmt.Println("this is receive method1", n)
}

func main() {
	structFieldMethod := new(StructFieldMethod)
	structFieldMethod.F1 = 1
	structFieldMethod.M1 = M1
	structFieldMethod.M1("m1")
	structFieldMethod.ReceiveMethod1("receive method1")
	fmt.Println(reflect.TypeOf(structFieldMethod.F1))
	fmt.Println(reflect.TypeOf(structFieldMethod.M1))
	fmt.Println(reflect.TypeOf(structFieldMethod.ReceiveMethod1))
}
