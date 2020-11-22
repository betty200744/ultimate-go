package encrypter_decrypter_cipher

import (
	"crypto/aes"
)

// Advanced Encryption Standard
// https://zh.wikipedia.org/wiki/%E9%AB%98%E7%BA%A7%E5%8A%A0%E5%AF%86%E6%A0%87%E5%87%86

const (
	Key = "exampleplainaaaa"
)

func AESEncrypt(plaintext string) []byte {
	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext must multiple of aes.BlockSize")
	}
	block, _ := aes.NewCipher([]byte(Key))
	ciphertext := make([]byte, len([]byte(plaintext)))
	block.Encrypt(ciphertext, []byte(plaintext))
	return ciphertext
}

func AESDecrypt(ciphertext []byte) string {
	block, _ := aes.NewCipher([]byte(Key))
	plaintext := make([]byte, len([]byte("exampleplaintext")))
	block.Decrypt(plaintext, ciphertext)
	return string(plaintext[:])
}
