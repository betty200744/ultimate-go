package encrypter_decrypter_cipher

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

var (
	IV     = []byte("exampleplainaaaa")
	Key, _ = hex.DecodeString("6368616e676520746869732070617373")
)

// Advanced Encryption Standard
// https://zh.wikipedia.org/wiki/%E9%AB%98%E7%BA%A7%E5%8A%A0%E5%AF%86%E6%A0%87%E5%87%86
func AESEncrypt(plaintext string) []byte {
	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext must multiple of aes.BlockSize")
	}
	block, _ := aes.NewCipher(Key)
	ciphertext := make([]byte, len([]byte(plaintext)))
	block.Encrypt(ciphertext, []byte(plaintext))
	return ciphertext
}
func AESDecrypt(ciphertext []byte) string {
	block, _ := aes.NewCipher(Key)
	plaintext := make([]byte, len(ciphertext))
	block.Decrypt(plaintext, ciphertext)
	return string(plaintext[:])
}

// https://zh.wikipedia.org/wiki/%E5%88%86%E7%BB%84%E5%AF%86%E7%A0%81%E5%B7%A5%E4%BD%9C%E6%A8%A1%E5%BC%8F
// CBC，Cipher-block chaining , like AESEncrypt
// https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation
func CBCEncrypt(plaintext []byte) []byte {
	block, _ := aes.NewCipher(Key)
	bs := block.BlockSize()
	padding := bs - len(plaintext)%bs
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	plaintext = append(plaintext, padText...)
	ciphertext := make([]byte, len(plaintext))
	modeE := cipher.NewCBCEncrypter(block, IV)
	modeE.CryptBlocks(ciphertext, plaintext)
	return ciphertext
}
func CBCDecrypt(ciphertext []byte) string {
	block, _ := aes.NewCipher(Key)
	modeD := cipher.NewCBCDecrypter(block, IV)
	plaintext := make([]byte, len(ciphertext))
	modeD.CryptBlocks(plaintext, ciphertext)
	length := len(plaintext)
	unpadding := int(plaintext[length-1])
	plaintext = plaintext[:(length - unpadding)]
	return string(plaintext)
}

// ECB, Electronic codebook
func ECBEncrypt(plaintext string) []byte {
	size := 16
	data := []byte(plaintext)
	ciphertext := make([]byte, len(data))
	block, _ := aes.NewCipher([]byte(Key))
	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Encrypt(ciphertext[bs:be], data[bs:be])
	}
	return ciphertext
}
func ECBDecrypt(ciphertext []byte) string {
	block, _ := aes.NewCipher([]byte(Key))
	plaintext := make([]byte, len(ciphertext))
	size := 16
	for bs, be := 0, size; bs < len(ciphertext); bs, be = bs+size, be+size {
		block.Decrypt(plaintext[bs:be], ciphertext[bs:be])
	}

	return string(plaintext[:])
}
