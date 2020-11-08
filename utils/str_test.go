package utils

import (
	"fmt"
	"testing"
)

func TestRandInt(t *testing.T) {
	fmt.Println(RandInt(1, 20))
}
func TestRandomString(t *testing.T) {
	fmt.Println(RandomString(5))
}
func TestToInt(t *testing.T) {
	fmt.Println([]byte(`0123456789`))
	fmt.Println([]byte(`ABCDEFGHIJKLMNOPQRSTUVWXYZ`))
	fmt.Println([]byte(`abcdefghijklmnopqrstuvwxyz`))
	a := "aaa"
	ToInt(a)
}
func TestString2Byte(t *testing.T) {
	b := []byte("aa")
	fmt.Println(b, string(b))
}
