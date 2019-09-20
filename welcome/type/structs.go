package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

type Shop struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
	Order []string
}

type S string

func main() {

	s := new(S)
	shop1 := new(Shop)
	fmt.Println("this shop1 id: ", shop1.Id == "", len(shop1.Owner) == 0)
	for _, value := range shop1.Owner {
		fmt.Println("this is shop1.Owner: ", value)
	}
	shopA := Shop{Name: "n1"}
	shopB := Shop{Name: "n2", Id: "b"}
	copier.Copy(shopA, shopB)
	fmt.Println("this is shopb: ", shopA)
	shop2 := Shop{}
	shopAP := &shopA
	fmt.Println(shopA)
	fmt.Println(&shopA)
	fmt.Println(shopA.Name)
	fmt.Println(*&shopA.Name)
	fmt.Println(shopAP)
	fmt.Println(shopAP.Name)
	fmt.Println(s)
	fmt.Println(shop2)
	fmt.Println(string(shop2.Id))
}
