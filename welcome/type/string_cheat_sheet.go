package main

import "fmt"

type Student struct {
	Id           int64
	Name         string
	Grade        int
	CurrentScore float32
}

func main() {
	student := Student{
		Id:           1,
		Name:         "betty",
		Grade:        5,
		CurrentScore: 3.3,
	}
	fmt.Println("print struct: ", fmt.Sprintf("%v", student))
	fmt.Println("print struct with field: ", fmt.Sprintf("%+v", student))
	fmt.Println("print struct type of value: ", fmt.Sprintf("%T", student))
	fmt.Println("print struct type and value: ", fmt.Sprintf("%#v", student))
}
