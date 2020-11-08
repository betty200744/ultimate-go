package encrypter_decrypter_cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"testing"
)

func Test_aes(t *testing.T) {
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	//Encryption
	plaintext := []byte("exampleplaintext")
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	modeE := cipher.NewCBCEncrypter(block, iv)
	modeE.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	fmt.Printf("%x\n", ciphertext)

	//Decryption
	iv = ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	// CBC modeE always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}
	modeD := cipher.NewCBCDecrypter(block, iv)
	// CryptBlocks can work in-place if the two arguments are the same.
	modeD.CryptBlocks(ciphertext, ciphertext)
	fmt.Printf("%s\n", ciphertext)
}
