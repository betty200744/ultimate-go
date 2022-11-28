package hello

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := Hello()
			fmt.Println(resp)
		})
	}
}
