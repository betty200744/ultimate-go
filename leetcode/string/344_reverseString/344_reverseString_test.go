package _44_reverseString

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	ReverseString([]byte(`hello`))
	ReverseString([]byte(`12345`))
}
