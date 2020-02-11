package main

import (
	"fmt"
	"github.com/jinzhu/copier"
	"reflect"
)

type Address struct {
	name    string
	street  string
	city    string
	state   string
	Pincode int
}

// Defining a struct type
type Shop struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Owner   string `json:"owner"`
	Order   []string
	Address Address `json:"address"`
}

type S string

func main() {
	// new a struct type
	s := new(S)
	shop1 := new(Shop)
	fmt.Println("this shop1 id: ", shop1.Id == "", len(shop1.Owner) == 0)
	for _, value := range shop1.Owner {
		fmt.Println("this is shop1.Owner: ", value)
	}
	// declaring and initializing a struct with name field
	shopA := Shop{Name: "n1"}
	shopB := Shop{Name: "n2", Id: "b"}
	// not name field
	shopC := Shop{"1", "n", "", []string{}, Address{}}
	copier.Copy(shopA, shopB)
	fmt.Println("this is shop b: ", shopB)
	shopAP := &shopA
	fmt.Println(shopA)
	fmt.Println(&shopA)
	// access fields of a struct
	fmt.Println(shopA.Name)
	fmt.Println(*&shopA.Name)
	fmt.Println(shopAP)
	fmt.Println(shopAP.Name)
	fmt.Println(s)
	// declaring and initializing a struct with zero value
	shop2 := Shop{}
	fmt.Println(shop2)
	fmt.Println(string(shop2.Id))
	fmt.Println(shopC)
	// iterate struct
	shop3 := Shop{Id: "id", Name: "n", Owner: "o", Order: []string{"1", "2"}, Address: Address{name: "add"}}
	shop3Value := reflect.ValueOf(shop3)
	shop3Values := make([]interface{}, shop3Value.NumField())
	for i := 0; i < shop3Value.NumField(); i++ {
		shop3Values[i] = shop3Value.Field(i).Interface()
	}
	fmt.Println("iterate struct: ", shop3Values)
}
