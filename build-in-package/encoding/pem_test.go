package encoding

import (
	"encoding/pem"
	"fmt"
	"testing"
)

func Test_pem_decode(t *testing.T) {
	block, _ := pem.Decode([]byte(ToPem(Key, "-----BEGIN RSA PRIVATE KEY-----\n", "-----END RSA PRIVATE KEY-----\n")))
	fmt.Println(block.Type)
	fmt.Println(block.Headers)
}
