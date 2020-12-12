package encrypter_decrypter_cipher

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestAES(t *testing.T) {
	ciphertext := AESEncrypt("exampleplaintext")
	plaintext := AESDecrypt(ciphertext)
	assert.Equal(t, plaintext, "exampleplaintext")
}
func TestCBC(t *testing.T) {
	ciphertext := CBCEncrypt([]byte("exampleplaintextexampleplaintext"))
	plaintext := CBCDecrypt(ciphertext)
	assert.Equal(t, plaintext, "exampleplaintextexampleplaintext")
}

func TestECB(t *testing.T) {
	ciphertext := ECBEncrypt("exampleplaintextexampleplaintext")
	plaintext := ECBDecrypt(ciphertext)
	assert.Equal(t, plaintext, "exampleplaintextexampleplaintext")
}
