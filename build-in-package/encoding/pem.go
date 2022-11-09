package encoding

import (
	"strings"
)

const (
	Key = ""
)

// pem , base64 ASCII, 用于x509 certificates
// pem, 常用后缀.crt, .pem, .cer, and .Key
func ToPem(key string, head, tail string) (pKey string) {
	var buffer strings.Builder
	buffer.WriteString(head)
	rawLen := 64
	keyLen := len(key)
	raws := keyLen / rawLen
	temp := keyLen % rawLen
	if temp > 0 {
		raws++
	}
	start := 0
	end := start + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(key[start:])
		} else {
			buffer.WriteString(key[start:end])
		}
		buffer.WriteByte('\n')
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString(tail)
	pKey = buffer.String()
	return
}
