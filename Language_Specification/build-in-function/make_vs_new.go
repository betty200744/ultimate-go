package main

import (
	"fmt"
)

/*
*  make, can only slice, map, channel type，方便append, delete等operator方法
*  make, return slice, map, or channel, not point，方便append, delete等operator方法
*  make, pass value, change a copy value, not self

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
	mSlice := make([]string, 0)
	nSlice := new([]string)
	mSlice = append(mSlice, "a", "b", "c")
	*nSlice = append(*nSlice, "a", "b", "c")
	fmt.Println("mSlice: ", mSlice)
	fmt.Println("nSlice: ", nSlice)

	mMap := make(map[string]string)
	mMap["m"] = "m"
	nMap := new(map[string]string)
	(*nMap)["n"] = "n"
	fmt.Println("mMap: ", mMap)
	fmt.Println("nMap: ", nMap)
	delete(mMap, "m")
	delete(*nMap, "n")
	fmt.Println("mMap, after delete: ", mMap)
	fmt.Println("nMap, after delete: ", nMap)

	chm := make(chan int)
	fmt.Println("chm: ", chm)

	// new
	sn := new([]string)
	chn := new(chan int)
	*sn = append(*sn, "a")
	fmt.Println("sn:", (*sn)[0])
	fmt.Println("chn:", chn)
}
