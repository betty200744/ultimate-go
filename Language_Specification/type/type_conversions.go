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
	x, y := 3, 4
	var f float64 = math.Sqrt(float64(x*x + y + y))
	var z uint = uint(f)
	fmt.Println(x, y, f, z)

	s := Status{
		Code:    1,
		Message: "ha ha",
	}
	fs := (FormatStatus)(s)
	fmt.Println(fmt.Sprintf("fs type is : %v", reflect.TypeOf(fs)))

}
