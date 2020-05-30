package utils

import (
	"math/rand"
)

func GetArrayOfLenAndSize(n, s int) []int {
	rand.Seed(86)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(10)
	}
	return arr
}
