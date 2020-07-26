package main

import "fmt"

/* syntax

defining:
// can accept both pointer and value receiver
// 如果receivers不是pointer， 那么是Type的Copy, 则无法change
func (r Type) functionName(...Type) Type {
    ...
}

call:
// can be call both pointer and value Type
Type.methodName(...)

 method:
	- is a function
    - different struct can have some method name, Decoupling即解偶的一种
	- nest method can promoted
 receiver type:
	- type and method must defined in some package, 只能是本package的struct， 即local type
	- 一般接收Pointer receivers ,如果receivers不是pointer， 那么是Type的Copy, 则无法change

When we declare a type, we must ask ourselves immediately:
 - Does this type require value semantic or pointer semantic?
 - If I need to modify this value:
	- should we create a new value
	- should we modify the value itself so everyone can see it?
 - It needs to be consistent. It is okay to guess it wrong the first time and refactor it later.
*/

// struct type
type MartReview struct {
	Id       string
	Name     string
	Merchant string
}

// method, value receiver
func (mr MartReview) FindById(id string) MartReview {
	fmt.Println(mr)
	return mr
}

// method, value receiver, 无法修改成功，因为是value receiver
func (mr MartReview) ChangeName() {
	mr.Name = "update name"
}
func (mr MartReview) ChangeMerchant() {
	mr.Merchant = "update merchant"
}

// method, pointer receiver
func (mr *MartReview) FindByIdTwo(id string) MartReview {
	return *mr
}

// method, pointer receiver
func (mr *MartReview) ChangeNameTwo() {
	mr.Name = "update name"
}
func (mr *MartReview) ChangeMerchantTwo() {
	mr.Merchant = "update merchant"
}
func main() {
	// Value Type, can call methods declared both value and pointer received
	mr := MartReview{Id: "1", Name: "mr", Merchant: "mer"}
	mr.FindById("1")
	mr.ChangeName()
	// Pointer Type, can call methods declared both value and pointer received
	mr2 := &MartReview{Id: "1", Name: "mr", Merchant: "mer"}
	mr2.FindById("1")
	mr2.ChangeName()

	fmt.Println("MartReview value type: ")
	mr.ChangeMerchant()
	fmt.Println("value receiver,  修改失败", mr.FindById("1"))
	mr.ChangeMerchantTwo()
	fmt.Println("pointer receiver, 修改成功 ", mr.FindByIdTwo("1"))

	fmt.Println("*MartReview pointer type: ")
	mr2.ChangeMerchant()
	fmt.Println("value receiver, 修改失败", mr2.FindById("1"))
	mr2.ChangeMerchantTwo()
	fmt.Println("pointer receiver, 修改成功", mr2.FindByIdTwo("1"))
}
