package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	fmt.Println("this is start")
	//log.Panic("this is panic")
	log.Print("this is print, ", "this is print 2")
	log.Println("this is printn, ", "this is printn2")
	err := errors.New("get env fail")
	log.Println(err)
	log.Fatal("this is fatal")
	fmt.Println("this is end")
}
