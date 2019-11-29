package main

import (
	"fmt"
)

type Parent3 struct {
	Age  string
	Name string
}

func main() {
	parents := make([]Parent3, 0)
	ids := []string{"4", "1", "5", "3", "2"} //期望结果
	p1 := Parent3{"3", "betty"}
	p2 := Parent3{"1", "cc"}
	p3 := Parent3{"2", "dd"}
	p4 := Parent3{"4", "ee"}
	p5 := Parent3{"5", "ff"}
	s := map[string]Parent3{}

	parents = append(parents, p1, p2, p3, p4, p5)
	for _, v := range parents {
		s[v.Age] = v
	}
	result := make([]Parent3, 0)
	for _, v := range ids {
		result = append(result, s[v])
	}

	fmt.Println(result)
}
