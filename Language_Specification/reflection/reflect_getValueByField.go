package main

import (
	"fmt"
	"reflect"
)

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

func main() {
	s1 := Student{1, 1, "Jack", 22}
	s2 := Student{2, 2, "betty", 22}
	ss := []*Student{&s1, &s2}
	ids := make([]int64, len(ss))
	ims := make([]int32, len(ss))
	for i, s := range ss {
		ids[i] = getValueByField("Id", s).(int64)
		ims[i] = getValueByField("ImagId", s).(int32)
	}
	fmt.Println(ids)
	fmt.Println(ims)
}
