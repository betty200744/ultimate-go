package encoding

import (
	"encoding/base64"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func Test_base64_EncodeToString(t *testing.T) {
	msg := "hello, world"
	base64msg := base64.StdEncoding.EncodeToString([]byte(msg))
	desMsg, _ := base64.StdEncoding.DecodeString(base64msg)
	assert.Equal(t, msg, string(desMsg))
	// echo c2t5aXRhY2hpY29kZXJAZ21haWwuY29tCg== | base64 -d
	email := "skyitachicoder@gmail.com"
	base64email := base64.StdEncoding.EncodeToString([]byte(email))
	desEmail, _ := base64.StdEncoding.DecodeString(base64email)
	assert.Equal(t, email, string(desEmail))
}
