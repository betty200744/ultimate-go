package _44_reverseString

import "fmt"

func ReverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	fmt.Println(string(s))
}
