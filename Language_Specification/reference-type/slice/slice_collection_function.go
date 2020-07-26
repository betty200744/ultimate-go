package main

import (
	"fmt"
	"reflect"
	"strings"
)

func GetItems(vs interface{}) []interface{} {
	v := reflect.ValueOf(vs)
	var items []interface{}
	for i := 0; i < v.Len(); i++ {
		items = append(items, v.Index(i).Interface())
	}
	return items
}
func Index(vs interface{}, t interface{}) int {
	items := GetItems(vs)
	for i, v := range items {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs interface{}, t interface{}) bool {
	return Index(vs, t) >= 0
}

func Any(vs interface{}, f func(interface{}) bool) bool {
	items := GetItems(vs)
	for _, v := range items {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs interface{}, f func(interface{}) bool) bool {
	items := GetItems(vs)
	for _, v := range items {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs interface{}, f func(interface{}) bool) []interface{} {
	vsf := make([]interface{}, 0)
	items := GetItems(vs)
	for _, v := range items {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs interface{}, f func(interface{}) interface{}) []interface{} {
	items := GetItems(vs)
	vsm := make([]interface{}, len(items))
	for i, v := range items {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {

	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v interface{}) bool {
		return strings.HasPrefix(v.(string), "p")
	}))

	fmt.Println(All(strs, func(v interface{}) bool {
		return strings.HasPrefix(v.(string), "p")
	}))

	fmt.Println(Filter(strs, func(v interface{}) bool {
		return strings.Contains(v.(string), "e")
	}))
}
