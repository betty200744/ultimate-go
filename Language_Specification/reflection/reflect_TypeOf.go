package main

import (
	"fmt"
	"reflect"
)

func main() {

	tst := "string"
	tst2 := 10
	tst3 := 1.2
	fmt.Println(reflect.TypeOf(tst))
	fmt.Println(reflect.TypeOf(tst2))
	fmt.Println(reflect.TypeOf(tst3))

}
