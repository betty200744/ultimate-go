package main

import (
	"fmt"
	"reflect"
)

/*
struct: field
struct: field function
struct: method: 如此一来我们变可以添加相同名字的method，因为receive不同
struct: type definition 与method definition需要在同一package
*/
type Range struct {
	From int64
	To   int64
}
type Owner struct {
	Id    int64
	Phone int64
}
type Company struct {
	A int32
	B int64
	C string
	Owner
	TagIds []string
	Price  Range
	F1     func(s string) string // field function 一般不这么做，不符合OOP的概念
}

func (s *Company) PrintCompany() { // struct method
	fmt.Println("this is receive method1", s)
}

func main() {
	// Initialize
	company := &Company{
		A: int32(32),
		B: int64(64),
		C: "skidoo",
		Owner: Owner{
			Id:    int64(333),
			Phone: int64(444),
		},
		TagIds: []string{"t1", "t2"},
		Price:  Range{From: int64(2), To: int64(3)},
		F1:     func(s string) string { return s },
	}
	// Selector and Promoted
	fmt.Println(company.A, company.Id, company.Phone)
	// iterate struct field name and value
	v := reflect.ValueOf(company).Elem()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		fmt.Printf("index : %d, name: %s , type: %s , value:  %v\n", i, f.Type().Name(), f.Type(), f.Interface())
	}
	// call function
	company.F1("m1")
	// call method
	company.PrintCompany()
}
