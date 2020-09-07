package _04_countPrimes

import (
	"fmt"
	"math"
)

func IsPrimes(n int) bool {
	if n < 2 {
		return false
	}
	if n != 2 && n%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func CountPrimes(n int) int {
	if n < 2 {
		return 0
	}
	arr := make([]int, 0)
	for i := 1; i < n; i++ {
		if IsPrimes(i) {
			arr = append(arr, i)
		}
	}
	fmt.Println(arr)
	return len(arr)
}
func GenePrimes(n int) []int {
	arr := make([]int, 0)
	for i := 1; i < n; i++ {
		if IsPrimes(i) {
			arr = append(arr, i)
		}
	}
	return arr
}
func GapPrimes(primes []int, k int) []int {
	gaps := make([]int, 0)
	for i := 0; i < len(primes)-1; i++ {
		gaps = append(gaps, primes[i+1]-primes[i])
	}
	arr := make([]int, 0)
	for i, i2 := range gaps {
		if i2 < k {
			arr = append(arr, i+1)
		}
	}
	return arr
}
