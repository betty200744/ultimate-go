package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type TTTT1 struct {
	X int     `json:"xx, ,string"` // json name: xx  , int to string
	Y int     `json:"yy"`
	U float32 `json:"uu"`
	A *[]int  `json:"aa"`
}

func (TTTT1) Do(event string) {
	fmt.Println(event)
}

type TTTT2 struct {
	x int `json:"x"`
	y int
	u float32
	A *[]int
}

type TTTTT struct {
	TTTT1
	TTTT2
}

type TT struct {
	f1     string `one:"1" two:"2" blank:""` // 3 tag
	f2     string ``
	f3     string `f three`
	f4, f5 int    `f four and five`
}

func (TT) Do(event string) {
	fmt.Println(event)
}

func main() {
	t1 := &TTTT1{X: 444, Y: 22, U: 3.3, A: &[]int{1, 2, 3}}
	b, _ := json.Marshal(t1)
	fmt.Println(string(b))

	t := reflect.TypeOf(TT{"1", "2", "3", 4, 5})
	fmt.Println(t.Name())
	fmt.Println(t.Field(0))
	fmt.Println(t.Field(0).Name)
	fmt.Println(t.FieldByName("f1"))

	method, _ := t.MethodByName("Do")
	fmt.Println(method.Name)
	f := method.Func
	in := make([]reflect.Value, 2)
	in[0] = reflect.ValueOf(TT.Do)
	in[1] = reflect.ValueOf("a")
	fmt.Println(f.Call(in))

	f1 := t.Field(0)
	fmt.Println(f1.Tag)
	f1v, _ := f1.Tag.Lookup("one")
	fmt.Println(f1v)

}
