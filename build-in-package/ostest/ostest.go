package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("DB", "product")
	fmt.Printf("env : %s , db:  %s.\n", os.Getenv("ENVIRON"), os.Getenv("DB"))
}
