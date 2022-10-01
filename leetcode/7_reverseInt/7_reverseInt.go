package __reverseInt

import (
	"math"
	"strconv"
)

func ReverseInt(x int) int {
	if x > 1<<31-1 || x < -(1<<31) {
		return 0
	}
	abs := math.Abs(float64(x))
	s := []byte(strconv.FormatInt(int64(abs), 10))
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	i, _ := strconv.Atoi(string(s))
	if i > 1<<31-1 || i < -(1<<31) {
		return 0
	}
	if x > 0 {
		return i
	} else {
		return -i
	}
}
