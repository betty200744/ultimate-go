package main

import "fmt"

type Merchant struct {
	Phone string `json:"phone"`
}

// struct type, like class
type MartReview struct {
	id   string
	name string
	Merchant
}

type Product struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

/* syntax
 method: is a function
 method: like class method , class is MartReview
 method: can accept both pointer and value receiver
 method: different struct can have some method name
 method: nest method can promoted
 receiver type:  type and method must defined in some package, 只能是本package的struct， 即local type
 receiver type:  you can use `type MyString string`
 receiver type:  一般接收Pointer receivers ,如果receivers不是pointer， 那么是Type的Copy, 则无法change
defining:
func (r Type) functionName(...Type) Type {
    ...
}

call:
Type.methodName(...)

*/
func (mr MartReview) FindById(id string) MartReview {
	return mr
}

func (mr MartReview) ChangeName() {
	mr.name = "update"
}

func (mr *MartReview) ChangePointName() {
	mr.name = "update"
}

func (mr Product) FindById(id string) Product {
	return mr
}

func (mr *MartReview) FindByName(name string) MartReview {
	return *mr
}

func (me *Merchant) GetMerchantPhone() string {
	return me.Phone
}

func main() {
	merchant := Merchant{Phone: "13486"}
	mr := MartReview{id: "1", name: "mr", Merchant: merchant}
	fmt.Println("value receiver, before change: get mr", mr.FindById("1"))
	fmt.Println("pointer receiver, before change , get mr", mr.FindByName("mr"))
	mr.ChangeName()
	fmt.Println("value receiver, after change: still get mr", mr.FindById("1"))
	mr.ChangePointName()
	fmt.Println("pointer receiver, after change , then get update ", mr.FindByName("update"))
	fmt.Println(merchant.GetMerchantPhone())
	fmt.Println(mr.GetMerchantPhone())
}
