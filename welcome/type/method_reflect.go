package main

import (
	"fmt"
	"reflect"
)

type T struct {
}

func (t T) M1(text string, number int) {

}
func (t *T) M2(map[string]int) {

}

func PrintMethodSet(value interface{}) {
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Printf("Method %d is : %s \n", i, t.Method(i).Name)
	}
}

func PrintFunction(val interface{}) {
	t := reflect.TypeOf(val)
	for i := 0; i < t.NumIn(); i++ {
		fmt.Printf("Funtion Name: %s, arg%d: %v \n", t, i, t.In(i))
	}
}

func main() {
	PrintFunction(T.M1)
	PrintFunction((*T).M2)
	PrintMethodSet(T{})
	PrintMethodSet(&T{})
}
