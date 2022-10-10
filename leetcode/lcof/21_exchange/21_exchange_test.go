package _1_exchange

import (
	"fmt"
	"testing"
)

//func TestRemove(t *testing.T) {
//	res := Remove([]int{2, 16, 3, 5, 13, 1, 16, 1, 12, 18, 11, 8, 11, 11, 5, 1}, 3)
//	fmt.Println(res)
//}

func TestExchange(t *testing.T) {
	fmt.Println(16 % 2)
	res := Exchange([]int{1, 2, 3, 4})
	fmt.Println(res)
}
func TestExchange2(t *testing.T) {
	res := Exchange([]int{2, 16, 3, 5, 13, 1, 16, 1, 12, 18, 11, 8, 11, 11, 5, 1})
	fmt.Println(res)
}
func TestExchange3(t *testing.T) {
	res := Exchange([]int{8, 10, 3, 20, 12, 4, 10, 8, 4, 0, 5, 17, 7, 20, 3})
	fmt.Println(res)
}
