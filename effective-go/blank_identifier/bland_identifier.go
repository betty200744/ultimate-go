package main

import (
	"fmt"
	_ "net/http/pprof"
	"os"
)

func main() {
	// the blank identifier in multi assignment
	if _, err := os.Stat("a"); os.IsNotExist(err) {
		fmt.Println(err)
	}

	// unused variables, will use later
	fd, err := os.Open("test.go")
	_, _ = fd, err

	// import for side effect
	// _ "net/http/pprof" , only run init()

	// interface checks, type assert
	var val interface{} = []string{"a"}
	if _, ok := val.([]string); ok {
		fmt.Printf("value %v of type %T implements json.Marshaler\n", val, val)
	}
}
