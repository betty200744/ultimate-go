package strings

import (
	"testing"
)

func TestReverseByWord(t *testing.T) {
	s := "one two three"
	ReverseByWord(s)
}
func TestReverseByCharacter(t *testing.T) {
	s := "abcdefg"
	ReverseByCharacter(s)
}
