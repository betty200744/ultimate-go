package reflection_ValueOf

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

type SS struct {
	Id             int64
	Name           string
	Price          int64
	RealPrice      int64
	Description    string
	Video          []string
	Imgs           []string
	ImgSnapshot100 string
	Score          int32
	CompanyId      string
	Brand          string
	Poster         string
	Weight         int32
	Barcode        string
	ClassId        int64
	RichText       string
	DiscountPrice  int64
	ActivityTag    []string
	ProductId      string
	CategoryId     int32
	MainCategory   int32
	MiddleCategory int32
	MinorCategory  int32
	SubBrandId     int64
	AboutBrand     string
}

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
	tn := reflect.ValueOf(n)
	fmt.Println("ValueOf: ", tn)
	fmt.Println("ValueOf Int: ", tn.Int())
}
func Rstring() {
	s := "this is string"
	ts := reflect.ValueOf(s)
	fmt.Println("ValueOf: ", ts)
	fmt.Println("ValueOf String: ", ts.String())
	fmt.Println("ValueOf Len: ", ts.Len())
}
func Rarray() {
	s := [3]int{1, 2, 4}
	st := reflect.ValueOf(s)
	fmt.Println(st.Elem())
	fmt.Println("Len: ", st.Len())
}
func Rmap() {
	m := map[string]string{"a": "a"}
	tm := reflect.ValueOf(m)
	fmt.Println("Typeof: ", tm)
	fmt.Println("Elem: ", tm.Elem())
}
func Rstruct() {
	t := T{F1: "vf1", F2: "vf2"}
	tt := reflect.ValueOf(t)
	val := reflect.ValueOf(t)
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
	tf := reflect.ValueOf(f)
	fmt.Println(tf)
}

type Student struct {
	Id     int64  `json:"id"`
	ImagId int32  `json:"imag_id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
}

func getValueByField(field string, s interface{}) interface{} {
	elem := reflect.ValueOf(s).Elem()
	value := elem.FieldByName(field)
	switch fmt.Sprint(value.Type()) {
	case "bool":
		return value.Bool()
	case "int32":
		return int32(value.Int())
	case "int64":
		return value.Int()
	case "string":
		return value.String()
	}
	return value.Int()
}

func test() {
	b := true
	s := ""
	n := 1
	f := 1.0
	a := []string{"foo", "bar", "baz"}
	ss := &SS{Id: 1, Name: "n", Price: 4}
	r := reflect.ValueOf(&ss).Elem()
	r.FieldByName("Name").SetString("AAA")
	fmt.Println(reflect.ValueOf(b).Kind())
	fmt.Println(reflect.ValueOf(s).Kind())
	fmt.Println(reflect.ValueOf(n).Kind())
	fmt.Println(reflect.ValueOf(f).Kind())
	fmt.Println(reflect.ValueOf(a).Index(0).Kind()) // For slices and strings

	structFieldMethod := new(StructFieldMethod)
	fmt.Println(reflect.ValueOf(structFieldMethod.F1).Kind())
	fmt.Println(reflect.ValueOf(structFieldMethod.M1).Kind())
	fmt.Println(reflect.ValueOf(structFieldMethod.ReceiveMethod1).Kind())

	fmt.Println(reflect.ValueOf(structFieldMethod.F1))
	fmt.Println(reflect.ValueOf(structFieldMethod.M1))
	fmt.Println(reflect.ValueOf(structFieldMethod.ReceiveMethod1))
}
