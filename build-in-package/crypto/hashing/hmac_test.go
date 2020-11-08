package hashing

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
)

func Test_hmac(t *testing.T) {
	key, _ := hex.DecodeString("63687373")
	in := []byte("abc")
	mac := hmac.New(sha256.New, key)
	mac.Write(in)
	out := mac.Sum(nil)
	fmt.Println(hmac.Equal(in, out))
}
