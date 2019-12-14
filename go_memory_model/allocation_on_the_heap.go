package main

import (
	"fmt"
	"unsafe"
)

type myStruct struct {
	myBool   bool    // 1 byte
	myFloat  float64 // 8 bytes
	myInt    int32   // 4 bytes
	myString string  // 16
}

type myStructOptimized struct {
	myFloat  float64 // 8 bytes
	myBool   bool    // 1 byte
	myInt    int32   // 4 bytes
	myString string  // 16 bytes
}

func main() {
	a := myStruct{}
	b := myStructOptimized{}
	fmt.Println(unsafe.Sizeof(a)) // 24 bytes
	fmt.Println(unsafe.Sizeof(b)) // ordered 16 bytes

}
