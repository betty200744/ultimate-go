package encoding

import (
	"encoding/hex"
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func Test_Hex(t *testing.T) {
	msg := "hello"
	hexMsg := hex.EncodeToString([]byte(msg))
	decMsg, _ := hex.DecodeString(hexMsg)
	assert.Equal(t, msg, string(decMsg))
}
func Test_Hex_Dump(t *testing.T) {
	msg := "hello"
	res := hex.Dump([]byte(msg))
	hexMsg := hex.EncodeToString([]byte(msg))

	fmt.Println(res)
	fmt.Println(hexMsg)
}
