package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Page  int      `json:"page"`
	Items []string `json:"items"`
}

func main() {
	// json for  boolean, int, float, string, slice, map, struct
	mapU := make(map[string]string)
	b1, _ := json.Marshal(true)
	i1, _ := json.Marshal(111)
	f1, _ := json.Marshal(3.2)
	s1, _ := json.Marshal("hello")
	sl1, _ := json.Marshal([]string{"a", "b", "c"})
	m1, _ := json.Marshal(map[string]string{"a": "a", "b": "b"})

	stru1, _ := json.Marshal(&Response{
		Page:  1,
		Items: []string{"a", "b", "c"},
	})
	fmt.Println(string(b1))
	fmt.Println(string(i1))
	fmt.Println(string(f1))
	fmt.Println(string(s1))
	fmt.Println(string(sl1))
	fmt.Println(string(m1))
	fmt.Println(string(stru1))
	fmt.Println(json.Unmarshal(m1, &mapU)) // 断言是否解析错误
}
