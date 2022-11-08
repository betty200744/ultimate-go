package encrypter_decrypter_cipher

import (
	"crypto/aes"
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func Test_Cipher(t *testing.T) {
	s := "fdaserearewrearf"
	b, _ := aes.NewCipher([]byte("qwertyuiopasdfgd"))
	in := []byte(s)
	encryS := make([]byte, len(in))
	b.Encrypt(encryS, in)
	fmt.Printf("%x \n", encryS)
	desD := make([]byte, len(in))
	b.Decrypt(desD, encryS)
	assert.Equal(t, s, string(desD))
}
