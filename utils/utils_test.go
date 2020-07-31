package main

import (
	"fmt"
	"testing"
)

func TestEscapeHttpHeader(t *testing.T) {
	fmt.Println(EscapeHttpHeader("https://stackoverflow.com/questions/2032149/optional-parameters"))
}
func TestGetIp(t *testing.T) {
	fmt.Println(GetIp())
}
func TestInArray(t *testing.T) {
	fmt.Println(InArray("a", []string{"a", "b"}))
}

func TestMd5(t *testing.T) {
	fmt.Println(Md5([]byte("abc")))
}
func TestRandInt(t *testing.T) {
	fmt.Println(RandInt(1, 20))
}
func TestRandomString(t *testing.T) {
	fmt.Println(RandomString(5))
}
