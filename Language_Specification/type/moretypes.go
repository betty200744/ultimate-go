package main

import (
	"fmt"
)

func computer(fn func(int, int) int, x int, y int) int {
	return fn(x, y)
}

func add_add(x, y int) int {
	return x + y
}

func main() {
	// pointers type
	i, j := 42, 2701
	p := &i
	fmt.Println(p)
	fmt.Println(*p)
	p = &j
	fmt.Println(p)
	fmt.Println(*p)
	*p = 3
	fmt.Println(*p)

	// struct , 有固定的key
	type Review struct {
		id   int
		name string
	}
	review := Review{1, "n"}
	fmt.Println(review)
	fmt.Println(review.id) // accessed struct field using dot
	review.id = 3          // access struct and change value
	fmt.Println(review.id)
	// pointers to struct
	vv := &review
	fmt.Println(vv)
	fmt.Println(vv.id)

	// struct literals , 同map
	v1 := Review{1, "n1"}
	v2 := Review{2, "n2"}
	v3 := Review{3, "n3"}
	v4 := Review{id: 1, name: "d"} //
	fmt.Println(v1, v2, v3, v4)

	// array
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a)
	b := [6]int{1, 2, 3, 4, 5, 6}
	reviews := [2]Review{{1, "r1"}}
	fmt.Println(b)
	fmt.Println(reviews)

	// slices , 同array.slice
	reviews1 := reviews[0:1] // slice reviews to two array
	fmt.Println(reviews1)

	// slices, like references to arrays , 即修改了slices也就修改了原array
	reviews1[0].id = 3
	fmt.Println(reviews)
	// slice, literals , 长度可以不指定
	names := []string{"a", "b", "c"}
	fmt.Println(names)
	// slice, len , cap
	fmt.Println(names, len(names), cap(names))
	// slice, create with make, fixed length
	titles := make([]int, 5)
	fmt.Println(titles)

	// multi array , 类型为[]string的array[]
	arrays := [][]string{
		[]string{"1", "11"},
		[]string{"2", "22"},
	}
	fmt.Println(arrays)
	// array append
	names = append(names, "betty")
	fmt.Println(names)
	// range
	for i, v := range reviews { // like reviews.map(i, v)
		fmt.Println("this is reviews", reviews)
		fmt.Println("this i v : ", i, v)
	}
	// range 2 , _ 即占位符
	for _, v := range reviews { // like reviews.map(i, v)
		fmt.Println("this is reviews", reviews)
		fmt.Println(v)
	}

	// map, 没有固定的key
	order := make(map[string]string)
	order["a"] = "a"
	order["b"] = "b"
	ordera, ok := order["a"]
	fmt.Println(ordera)
	fmt.Println(ok)
	fmt.Println(order)
	user := map[string]string{"id": "1", "name": "n1"}
	fmt.Println("this is user:", user)
	// mutating maps
	sku := map[string]int{"id": 1, "price": 3}
	fmt.Println("this is sku:", sku)
}
