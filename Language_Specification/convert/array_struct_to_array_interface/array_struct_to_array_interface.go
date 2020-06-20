package main

import (
	"fmt"
	"reflect"
)

type PairKeyValue struct {
	Id    string
	Key   string
	Value int
}
type PairKeyValue2 struct {
	Id    int64
	Key   string
	Value int
}

func GetStringIds(input interface{}) []string {
	object := reflect.ValueOf(input)
	var items []interface{}
	for i := 0; i < object.Len(); i++ {
		items = append(items, object.Index(i).Interface())
	}

	// Populate the rest of the items into <ids>
	var ids []string
	for _, v := range items {
		val := reflect.ValueOf(v)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		id := val.FieldByName("Id").String()
		ids = append(ids, id)
	}
	return ids
}
func GetInt64Ids(input interface{}) []int64 {
	object := reflect.ValueOf(input)
	var items []interface{}
	for i := 0; i < object.Len(); i++ {
		items = append(items, object.Index(i).Interface())
	}
	// Populate the rest of the items into <ids>
	var ids []int64
	for _, v := range items {
		val := reflect.ValueOf(v)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		id := val.FieldByName("Id").Interface()
		ids = append(ids, id.(int64))
	}
	return ids
}
func main() {
	items := []*PairKeyValue{{Id: "1"}, {Id: "2"}}
	items2 := []PairKeyValue2{{Id: 1}, {Id: 2}}
	interfaces := GetStringIds(items)
	interfaces2 := GetInt64Ids(items2)
	fmt.Println(interfaces)
	fmt.Println(interfaces2)
}
