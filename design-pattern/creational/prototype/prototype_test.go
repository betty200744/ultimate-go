package prototype

import (
	"fmt"
	"testing"
)

func TestNewConfig(t *testing.T) {
	c := NewConfig("/tmp", "betty")
	fmt.Println(c)
}
func TestConfig_WithUser(t *testing.T) {
	c := NewConfig("", "")
	c.WithUser("betty")
	fmt.Println(c)
}
func TestConfig_WithWorkDir(t *testing.T) {
	c := NewConfig("", "")
	c.WithWorkDir("/home")
	fmt.Println(c)
}
