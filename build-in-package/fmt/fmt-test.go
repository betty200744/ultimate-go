package main

import "fmt"

func ErrorTest() error {
	e := fmt.Errorf(" %s%d found ? %t , this is chan %p , this is chan value %s", "sku ", 123, false)
	return e
}

func main() {

	/*
		basic
		%v	the value in a default format
	*/
	ch := make(chan int)
	s := fmt.Sprintf("bool : %v , int: %v , string: %v, chan: %v, point: %v", true, 1, "a", ch, &ch)
	fmt.Println(s)
	/*	default format

		bool:                    %t
		int, int8 etc.:          %d
		uint, uint8 etc.:        %d, %#x if printed with %#v
		float32, complex64, etc: %g
		string:                  %s
		chan:                    %p
		pointer:                 %p
	*/
	ch1 := make(chan int)
	s1 := fmt.Sprintf("bool : %t , int: %d , string: %s, chan: %p, point: %p", true, 1, "a", ch1, &ch1)
	fmt.Println(s1)

	/*
	   err
	*/

	fmt.Println(ErrorTest())

}
