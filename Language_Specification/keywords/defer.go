package main

import "fmt"

func Foo() {
	defer finished()
	defer func() {
		fmt.Println("defer 1")
	}()
	defer func() {
		fmt.Println("defer 2")
	}()
	fmt.Println("start")
	if true {
		return
	}
	nums := []int{1, 3, 4, 5}
	for key, value := range nums {
		fmt.Println(key, value)
	}
}

func finished() {
	fmt.Println("defer finished")
}

// defer倒序执行
func main() {
	Foo()
}
