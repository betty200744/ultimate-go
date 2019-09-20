package main

import (
	"fmt"
	"reflect"
)

func InArray(val interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {

				return true
			}
		}
	}

	return false
}

func main() {
	ll := []int64{1, 2, 3, 4}
	fmt.Println(InArray(int64(3), ll))
	fmt.Println(fmt.Sprintf("%016x%016x", 7285711369521225825, 5914568201030938479))
}
