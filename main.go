package main

import (
	"fmt"
)

type FooInterface interface {
	bar() // method receiver is value then interface is value , method receiver is pointer then interface is pointer
}

//Won't throw a compile error. Because the bar method belongs to Foo struct.
type Foo1 struct {
}

func (f Foo1) bar() {
	panic("implement me")
}

//Won't throw a compile error. Because the bar method belongs to Foo struct.
//Will throw a error. Because the bar method is bind to &Foo. And when compile checks whether Foo confirms to Foo interface, it will only check the required method existence.
type Foo2 struct {
}

func (f *Foo2) bar() {
	panic("implement me")
}

func funcName(foo FooInterface) {
	fmt.Println(foo)
}

func main() {
	funcName(Foo1{})  // Foo1.bar receiver is value , then Foo1 need value
	funcName(&Foo2{}) // Foo2.bar receiver is pointer, then Foo2 need pointer
}
