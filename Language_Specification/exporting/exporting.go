package main

import (
	"fmt"
	"ultimate-go/Language_Specification/exporting/codes"
)

// -------------------
// Exported identifier
// -------------------
func main() {
	// Exported identifier
	fmt.Println(codes.ArchiveAlreadyDel)
	// Access a value of an unexported identifier
	err100 := codes.New(100)
	fmt.Println(err100)
	// Unexported fields from an exported struct
	status := codes.Status{Name: "s"}
	fmt.Println(status)
}
