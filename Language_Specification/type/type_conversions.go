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
	fs := (FormatStatus)(s)
	fmt.Println(x, y, f, z)
	fmt.Println(fmt.Sprintf("s type is : %v", reflect.TypeOf(s)))
	fmt.Println(fmt.Sprintf("fs type is : %v", reflect.TypeOf(fs)))
}
