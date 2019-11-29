package main

import "fmt"
import "reflect"

type TC struct {
	Name string
}

func (t *TC) Foo() {
	fmt.Println("foo")
}

type YourT1 struct{}

func (y YourT1) MethodBar() {
	fmt.Println("no args")
}

type YourT2 struct{}

func (y YourT2) MethodFoo(i int, oo string) string {
	fmt.Printf("arg i: %d, arg oo: %s \n", i, oo)
	return oo
}

func Invoke(any interface{}, name string, args ...interface{}) reflect.Value {
	in := make([]reflect.Value, len(args))
	for i, _ := range args {
		in[i] = reflect.ValueOf(args[i])
	}

	// when use valueof , then method Func.call , in is args
	elemv := reflect.ValueOf(any)
	fmt.Println("this is elemv: ", elemv)
	methodv := elemv.MethodByName(name)
	fmt.Println("this is methodv: ", methodv.String())
	fmt.Println(methodv.Call(in)[0])

	// when use typeof , then method Func.call , first arg is receive, then args
	elemt := reflect.TypeOf(any)
	fmt.Println("this is elemt: ", elemt)
	m, _ := elemt.MethodByName(name)
	methodt := m.Func
	in = append([]reflect.Value{reflect.ValueOf(any)}, in...) // in should include receive
	fmt.Println("this is methodt:", methodt.String())
	fmt.Println(methodt.Call(in)[0])
	return methodt.Call(in)[0]
}

func main() {
	var t TC
	reflect.ValueOf(&t).MethodByName("Foo").Call([]reflect.Value{})

	a := Invoke(YourT2{}, "MethodFoo", 10, "abc")
	fmt.Println(a)
	//Invoke(YourT1{}, "MethodBar")
}
