package main

import (
	"fmt"
	"math"
	"reflect"
)

type Status struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}
type FormatStatus Status

func main() {
	x := 3
	y := 4
	f := math.Sqrt(float64(x*x + y + y))
	z := uint(f)
	s := Status{
		Code:    1,
		Message: "ha ha",
	}
	fs := FormatStatus(s)
	fmt.Println(x, y, f, z)
	fmt.Println(fmt.Sprintf("s type is : %v", reflect.TypeOf(s)))
	fmt.Println(fmt.Sprintf("fs type is : %v", reflect.TypeOf(fs)))

	// Conversions
	aaa := int32(10)
	fmt.Printf("aaa : %T [%v]", aaa, aaa)
	fmt.Printf(" float32(2.718281828): %T, [%v] \n", float32(2.718281828), float32(2.718281828))
	fmt.Printf(" complex128(1): %T, [%v] \n", complex128(1), complex128(1))
	fmt.Printf(" float32(0.49999999): %T, [%v] \n", float32(0.49999999), float32(0.49999999))
	fmt.Printf(" float64(-1e-1000): %T, [%v] \n", float64(-1e-1000), float64(-1e-1000))
	fmt.Printf(" string('x'): %T, [%v] \n", string('x'), string('x'))
	fmt.Printf(" string(0x266c): %T, [%v] \n", string(0x266c), string(0x266c))
	fmt.Printf(" string([]byte{'a'}): %T, [%v] \n", string([]byte{'a'}), string([]byte{'a'}))
	fmt.Printf(" (*int)(nil): %T, [%v] \n", (*int)(nil), (*int)(nil))
	type Person struct {
		Name    string
		Address *struct {
			Street string
			City   string
		}
	}

	var data *struct {
		Name    string `json:"name"`
		Address *struct {
			Street string `json:"street"`
			City   string `json:"city"`
		} `json:"address"`
	}
	var person = (*Person)(data)
	fmt.Printf("person: %T \n", person)
}
