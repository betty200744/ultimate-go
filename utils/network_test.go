package utils

import (
	"fmt"
	"testing"
)

func TestEscapeHttpHeader(t *testing.T) {
	fmt.Println(EscapeHttpHeader("https://stackoverflow.com/questions/2032149/optional-parameters"))
}
func TestGetIp(t *testing.T) {
	fmt.Println(GetIp())
}
