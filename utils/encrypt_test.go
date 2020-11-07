package utils

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestMd5Hex(t *testing.T) {
	s1 := Md5Hex([]byte(`aaa`))
	s2 := Md5Hex1([]byte(`aaa`))
	assert.Equal(t, s1, s2)
}

func TestMd5Hex1(t *testing.T) {
	s1 := Md5Hex([]byte(`aaa`))
	s2 := Md5Hex1([]byte(`aaa`))
	assert.Equal(t, s1, s2)
	fmt.Println(s2)
}

func TestMd5Base64(t *testing.T) {
	s1 := Md5Base64([]byte(`aaa`))
	fmt.Println(s1)
}
