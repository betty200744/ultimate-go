package rand

import (
	"crypto/rand"
	"fmt"
	"io"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("", func(t *testing.T) {
		res := make([]byte, 10)
		rand.Read(res)
		fmt.Println(string(res[:]))
	})
	t.Run("", func(t *testing.T) {
		tmpBytes := make([]byte, 16)
		io.ReadFull(rand.Reader, tmpBytes)
		fmt.Println(string(tmpBytes[:]))
	})
}
