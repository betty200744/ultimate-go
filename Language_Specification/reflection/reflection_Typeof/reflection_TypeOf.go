package reflection_Typeof

import (
	"fmt"
	"reflect"
)

type SomeStruct struct {
	SomeField uint32
}
type T struct {
	F1 string `json:"f1"`
	F2 string `json:"f2"`
}

func (t *T) name() {
	fmt.Println("name")
}

func Rint() {
	n := 1
	tn := reflect.TypeOf(n)
	fmt.Println("Typeof: ", tn)
}
func Rstring() {
	s := "this is string"
	ts := reflect.TypeOf(s)
	fmt.Println("Typeof: ", ts)
}
func Rarray() {
	s := [3]int{1, 2, 4}
	st := reflect.TypeOf(s)
	fmt.Println(st.Elem())
	fmt.Println("Len: ", st.Len())
}
func Rmap() {
	m := map[string]string{"a": "a"}
	tm := reflect.TypeOf(m)
	fmt.Println("Typeof: ", tm)
	fmt.Println("Key: ", tm.Key())
	fmt.Println("Elem: ", tm.Elem())
}
func Rstruct() {
	t := T{F1: "vf1", F2: "vf2"}
	tt := reflect.TypeOf(t)
	val := reflect.ValueOf(t)
	fmt.Println("Align: ", tt.Align())
	fmt.Println("FieldAlign: ", tt.FieldAlign())
	fmt.Println("Field:", tt.Field(0))
	fmt.Println("FieldByIndex: ", tt.FieldByIndex([]int{1}))
	fmt.Println(tt.FieldByName("F1"))
	fmt.Println("NumField: ", tt.NumField())
	fmt.Println("NumMethod: ", tt.NumMethod())
	fmt.Println(tt.MethodByName("name"))
	for i := 0; i < tt.NumField(); i++ {
		fieldValue := val.Field(i)
		fmt.Println(fieldValue)
	}
}
func Rfun() {
	f := func(a, b string) (c string) {
		c = "aaa"
		return c
	}
	tf := reflect.TypeOf(f)
	fmt.Println("NumIn: ", tf.NumIn())
	fmt.Println("NumIn: ", tf.NumOut())
	fmt.Println("NumIn: ", tf.In(0))
	fmt.Println("NumIn: ", tf.Out(0))
}
