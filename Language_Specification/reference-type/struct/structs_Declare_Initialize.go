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
	type User struct{ Name string }
	type User1 User   // new type User1
	type User2 = User // User2 is an alias of User
	user0 := User{Name: "user0"}
	user1 := User1{Name: "user1"}
	user2 := User2{Name: "user2"}
	fmt.Printf("user0 type: %T, value: %v\n", user0, user0)
	fmt.Printf("user1 type: %T, value: %v\n", user1, user1)
	fmt.Printf("user2 type: %T, value: %v\n", user2, user2)
	// Defining Name type
	type Shop struct {
		Id      string `json:"id"`
		Name    string `json:"name"`
		Owner   string `json:"owner"`
		Order   []string
		Address Address `json:"address"`
	}
	// Defining Anonymous type
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
