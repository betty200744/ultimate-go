package encoding

import (
	"encoding/hex"
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestHex(t *testing.T) {
	src := []byte("hello")
	encodedStr := hex.EncodeToString(src)
	fmt.Println(encodedStr)
	decodeStr, err := hex.DecodeString(encodedStr)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, src, decodeStr)
}
