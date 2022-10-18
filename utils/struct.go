package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func StructToMap(in interface{}) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("StructToMap only accepts structs; got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := typ.Field(i)
		if name := fi.Name; name != "" {
			out[name] = v.Field(i).Interface()
		}
	}
	return out, nil
}

func Struct2JsonMap(obj interface{}) map[string]interface{} {
	var data = make(map[string]interface{})
	structByte, err := json.Marshal(obj)
	if err != nil {
		return nil
	}
	_ = json.Unmarshal(structByte, &data)
	return data
}

// input like []*Shop or []Shop
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
