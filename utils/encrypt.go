package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

func Md5Hex(v []byte) string {
	h := md5.New()
	h.Write(v)
	s1 := h.Sum(nil)
	return hex.EncodeToString(s1)
}
func Md5Hex1(v []byte) string {
	hashInBytes := md5.Sum(v)
	return hex.EncodeToString(hashInBytes[:])
}

func Md5Base64(v []byte) string {
	h := md5.New()
	h.Write(v)
	s1 := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(s1)
}

func Sha1(input []byte) string {
	hash := sha1.New()
	hash.Write(input)
	return hex.EncodeToString(hash.Sum(nil))
}

func Sha256(input []byte) string {
	hash := sha256.New()
	hash.Write(input)
	return hex.EncodeToString(hash.Sum(nil))
}

func Sha512(input []byte) string {
	hash := sha512.New()
	hash.Write(input)
	return hex.EncodeToString(hash.Sum(nil))
}
