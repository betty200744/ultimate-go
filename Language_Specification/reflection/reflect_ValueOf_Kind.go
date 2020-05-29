package main

import (
	"fmt"
	"reflect"
)

type FieldFunction func(s string) string
type StructFieldMethod struct {
	F1 int64
	F2 string
	M1 FieldFunction // 为struct添加method的方式一， 一般不这么做，不符合OOP的概念
}

func (s *StructFieldMethod) ReceiveMethod1(n string) { // 为struct添加method的方式二
	fmt.Println("this is receive method1", n)
}

func main() {
	b := true
	s := ""
	n := 1
	f := 1.0
	a := []string{"foo", "bar", "baz"}

	fmt.Println(reflect.ValueOf(b).Kind())
	fmt.Println(reflect.ValueOf(s).Kind())
	fmt.Println(reflect.ValueOf(n).Kind())
	fmt.Println(reflect.ValueOf(f).Kind())
	fmt.Println(reflect.ValueOf(a).Index(0).Kind()) // For slices and strings

	structFieldMethod := new(StructFieldMethod)
	fmt.Println(reflect.ValueOf(structFieldMethod.F1).Kind())
	fmt.Println(reflect.ValueOf(structFieldMethod.M1).Kind())
	fmt.Println(reflect.ValueOf(structFieldMethod.ReceiveMethod1).Kind())

	fmt.Println(reflect.TypeOf(structFieldMethod.F1))
	fmt.Println(reflect.TypeOf(structFieldMethod.M1))
	fmt.Println(reflect.TypeOf(structFieldMethod.ReceiveMethod1))
}
