package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
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

func Test_bcrypt(t *testing.T) {
	passwork := "a"
	b, _ := bcrypt.GenerateFromPassword([]byte("a"), 1)
	fmt.Println(string(b))
	fmt.Println(bcrypt.Cost(b))
	err := bcrypt.CompareHashAndPassword(b, []byte(passwork))
	assert.Empty(t, err)
}

func TestSha1(t *testing.T) {
	s := Sha1([]byte("a"))
	fmt.Println(s)
}

func TestSha256(t *testing.T) {
	s := Sha256([]byte("a"))
	fmt.Println(s)
}

func TestSha512(t *testing.T) {
	s := Sha512([]byte("a"))
	fmt.Println(s)
}
