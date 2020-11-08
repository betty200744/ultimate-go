package hashing

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

func EncodeMD5Hex(v []byte) string {
	m := md5.New()
	m.Write(v)
	return hex.EncodeToString(m.Sum(nil))
}

func EncodeMD5Base64(v []byte) string {
	m := md5.New()
	m.Write(v)
	return base64.StdEncoding.EncodeToString(m.Sum(nil))
}

func Md5(v []byte) []byte {
	h := md5.New()
	h.Write([]byte("abc"))
	return h.Sum(nil)
}
