package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := []string{"a", "b"}
	b := []int32{1}
	i := make([]int32, 3)
	copy(i, b)
	for _, value := range a {
		value = fmt.Sprintf("%s:%s", "prefix", value)
	}
	fmt.Println(i[0])
	fmt.Println(i[1])
	fmt.Println(i[2])
	fmt.Println(strconv.FormatInt(int64(0), 10))
}
