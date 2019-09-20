package main

import (
	"fmt"
)

/*
*  make, can only slice, map, channel type
*  make, return slice, map, or channel, not point
*  make, pass value, change a copy value, not self
*  make, 以便map用make

Call             Type T     Result
make(T, n)       slice      slice of type T with length n and capacity n
make(T, n, m)    slice      slice of type T with length n and capacity m

make(T)          map        map of type T
make(T, n)       map        map of type T with initial space for approximately n elements

make(T)          channel    unbuffered channel of type T
make(T, n)       channel    buffered channel of type T, buffer size n

*/

/*
* new, can any Type
* new, return *Type
* new, pass point, change self
* new, 一般struct用new

new(T) , init length is 0 , cannot be index,
*/

func main() {
	// make
	sm := make([]string, 0)
	mapm := make(map[string]string)
	chm := make(chan int)
	sm = append(sm, "a")
	mapm["a"] = "a"
	fmt.Println("sm: ", sm)
	fmt.Println("mapm: ", mapm)
	fmt.Println("chm: ", chm)

	// new
	sn := new([]string)
	chn := new(chan int)
	*sn = append(*sn, "a")
	fmt.Println("sn:", (*sn)[0])
	fmt.Println("chn:", chn)
}
