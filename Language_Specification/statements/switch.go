package main

import (
	"fmt"
	"runtime"
	"time"
)

func Do(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Println("i am string")
	case int:
		fmt.Println("i am int")
	case bool:
		fmt.Println("i am bool")
	default:
		fmt.Println("i don't know")
	}
}

func Switch() {
	// switch variable
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s", os)
	}
	// switch true
	h := time.Now().Hour()
	switch {
	case h < 12:
		fmt.Println("good morning")
	case h < 17:
		fmt.Println("good afternoon")
	default:
		fmt.Println("good eventing")
	}

	// value switch
	Do("string")
	Do(123)
	Do(true)
}
