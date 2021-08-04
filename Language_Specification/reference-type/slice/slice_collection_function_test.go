package slice

import (
	"fmt"
	"strings"
	"testing"
)

func TestAll(t *testing.T) {
	type args struct {
		vs interface{}
		f  func(interface{}) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var strs = []string{"peach", "apple", "pear", "plum"}

			fmt.Println(Index(strs, "pear"))

			fmt.Println(Include(strs, "grape"))

			fmt.Println(Any(strs, func(v interface{}) bool {
				return strings.HasPrefix(v.(string), "p")
			}))

			fmt.Println(All(strs, func(v interface{}) bool {
				return strings.HasPrefix(v.(string), "p")
			}))

			fmt.Println(Filter(strs, func(v interface{}) bool {
				return strings.Contains(v.(string), "e")
			}))

			if got := All(tt.args.vs, tt.args.f); got != tt.want {
				t.Logf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}
