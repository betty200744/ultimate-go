package main

import (
	"fmt"
)

type Address struct {
	name    string
	street  string
	city    string
	state   string
	Pincode int
}

func main() {
	// Defining Name type, 相同结构不可=
	type Shop struct {
		Id      string `json:"id"`
		Name    string `json:"name"`
		Owner   string `json:"owner"`
		Order   []string
		Address Address `json:"address"`
	}
	// Defining Anonymous type， 相同结构可=
	shopAnonymous := struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Owner string `json:"owner"`
	}{Id: "1", Name: "Anonymous Type"}
	fmt.Printf("shopAnonymous: %T, [%v] \n", shopAnonymous, shopAnonymous)

	// initializing a struct with zero value
	shop0 := Shop{}
	fmt.Println("initializing a struct with zero value: ", shop0)

	// initializing a struct with name field and access fields of a struct
	shop1 := Shop{Name: "n1", Id: "b"}
	fmt.Println("initializing a struct with name field: ", shop1)
	fmt.Println("access fields of a struct", shop1.Name)

	// initializing a struct with not name field
	shop2 := Shop{"1", "n", "", []string{}, Address{}}
	fmt.Println("initializing a struct with not name field: ", shop2)

	// new a struct type with zero value
	shop3 := new(Shop)
	fmt.Println("new a struct type with zero value: ", shop3)
}
