package encrypter_decrypter_cipher

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestAESEncrypt(t *testing.T) {
	ciphertext := AESEncrypt("exampleplaintext")
	fmt.Printf("%x \n", ciphertext)
}

func TestAESDecrypt(t *testing.T) {
	ciphertext := AESEncrypt("exampleplaintext")
	plaintext := AESDecrypt(ciphertext)
	assert.Equal(t, plaintext, "exampleplaintext")
}
