package utils

import (
	"crypto/md5"
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
