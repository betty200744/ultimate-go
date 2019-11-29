package main

import (
	"fmt"
	"reflect"
)

type T struct {
	F1 string `json:"f1"`
	F2 string `json:"f2"`
}

func main() {
	n := 1
	s := "this is string"
	t := T{F1: "vf1", F2: "vf2"}
	m := map[string]string{"a": "a"}

	tn := reflect.TypeOf(n)
	ts := reflect.TypeOf(s)
	tt := reflect.TypeOf(t)
	tm := reflect.TypeOf(m)
	ttf1, ttf1pre := tt.FieldByName("F1")
	//f1 {name, pkgpath, type, tag, offset, index, present}
	fmt.Println(tn)
	fmt.Println(ts)
	fmt.Println(tt)
	fmt.Println(ttf1, ttf1pre)
	fmt.Println(tm)
}
