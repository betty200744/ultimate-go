package main

import (
	"fmt"
	"strings"
)

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func StringEscape() {
	s := "a string ++"
	fmt.Println("s: ", s)

	// Trim one trailing '+'.
	s1 := s
	if last := len(s1) - 1; s1[last] == '+' {
		s1 = s1[:last]
	}
	fmt.Println("s1:", s1)

	// Trim all trailing '+'.
	s2 := s
	s2 = strings.TrimRight(s2, "+")
	fmt.Println("s2:", s2)

	// Trim suffix "+".
	s3 := s
	s3 = TrimSuffix(s3, "+")
	fmt.Println("s3:", s3)
}
