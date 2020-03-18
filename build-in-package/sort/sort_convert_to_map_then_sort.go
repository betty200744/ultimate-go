package main

import (
	"fmt"
)

type Parent3 struct {
	Age  string
	Name string
}

func main() {
	ids := []string{"4", "1", "5", "3", "2"} // 参考排序
	parents := []Parent3{                    // 待排序列表
		{"3", "betty"},
		{"1", "cc"},
		{"2", "dd"},
		{"4", "ee"},
		{"5", "ff"},
	}
	parentsMap := map[string]Parent3{}
	for _, v := range parents {
		parentsMap[v.Age] = v
	}
	result := make([]Parent3, 0) // 排序后结果
	for _, v := range ids {
		result = append(result, parentsMap[v])
	}

	fmt.Println(result)
}
