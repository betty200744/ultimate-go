package simple_factory

import (
	"fmt"
	"testing"
)

func TestNewIdGenerator(t *testing.T) {
	s, _ := NewIdGenerator()
	id := s.Generate()
	fmt.Println(id)
}
