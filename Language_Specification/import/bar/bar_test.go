package bar

import (
	"fmt"
	"testing"

	"ultimate-go/Language_Specification/import/foo"
)

func TestBar_Bar(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &foo.Foo{}
			b := &Bar{}
			fmt.Println(b, f)
		})
	}
}
