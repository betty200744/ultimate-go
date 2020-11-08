package hashing

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMd5(t *testing.T) {
	Md5([]byte(`aaa`))
}

func TestEncodeMD5Hex(t *testing.T) {
	res := EncodeMD5Hex([]byte(`aaa`))
	assert.Equal(t, "47bce5c74f589f4867dbd57e9ca9f808", res)
	fmt.Println(res)
}

func TestEncodeMD5Base64(t *testing.T) {
	m := Md5([]byte(`aaa`))
	fmt.Println(string(m))
	res := EncodeMD5Base64([]byte(`aaa`))
	fmt.Println(res)
}
