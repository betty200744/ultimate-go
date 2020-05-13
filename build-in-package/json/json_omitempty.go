package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Dummy struct {
	Name    string  `json:"name"`
	Number  int64   `json:"number"`
	Pointer *string `json:"pointer"`
}

func main() {
	pointer := "yes"

	dummyComplete := Dummy{
		Name:    "Mr Dummy",
		Number:  4,
		Pointer: &pointer,
	}

	data, err := json.Marshal(dummyComplete)
	if err != nil {
		fmt.Println("An error occured: %v", err)
		os.Exit(1)
	}

	fmt.Println(string(data))

	// ALL of those are considered empty by Go
	dummyNilPointer := Dummy{
		Name:    "",
		Number:  0,
		Pointer: nil,
	}

	dataNil, err := json.Marshal(dummyNilPointer)
	if err != nil {
		fmt.Println("An error occured: %v", err)
		os.Exit(1)
	}

	fmt.Println(string(dataNil))
}
