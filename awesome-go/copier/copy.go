package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

type User struct {
	Name string
	Age  int32
	Role string
}

type Employee struct {
	Name      string
	Age       int32
	DoubleAge int32
	EmployeId int64
	SuperRole string
}

func (user *User) DoubleAge() int32 {
	return 2 * user.Age
}

func (employee *Employee) Role(role string) {
	employee.SuperRole = "Super " + role
}

func main() {
	user := &User{Name: "Jinzhu", Age: 18, Role: "Admin"}
	employee := &Employee{}

	fmt.Printf("user type:  %T \n", user)
	fmt.Printf("employee type: %T \n", employee)

	copier.Copy(employee, user)

	fmt.Println("this is user: ", user)
	fmt.Println("this is copy employee: ", employee)

	users := &[]User{{Name: "Jinzhu", Age: 18, Role: "Admin"}, {Name: "jinzhu 2", Age: 30, Role: "Dev"}}
	employees := &[]Employee{}
	// Copy struct to struct
	copier.Copy(&employees, &user)

	fmt.Println(*employee)

	// Copy slice to slice
	employees = &[]Employee{}
	copier.Copy(&employees, &users)

	a := struct {
		Id string `json:"id"`
	}{Id: "a"}

	b := struct {
		Id string `json:"id"`
	}{Id: "b"}

	c := struct {
		Id string `json:"id"`
	}{}
	_ = copier.Copy(&c, &a)
	_ = copier.Copy(&c, &b)

	fmt.Println("i am c : ", c)
}
