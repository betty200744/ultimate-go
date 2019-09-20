package main

import (
	"fmt"
	"os"
)

func main() {
	// good
	if 1 < 2 {
		fmt.Println(true)
	}

	if err := os.Chmod("./effective-go/control_structures/if.go", 0664); err != nil {
		fmt.Println(err.Error())
	}

	f, err := os.Open("./effective-go/control_structures/test")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(f.Read([]byte("s")))
}
