package hashset

import (
	"fmt"
	"testing"
)

func Test_Hashset(t *testing.T) {
	hashSet := New("a", "b")
	hashSet.Add("c")
	hashSet.Add("d")
	fmt.Println(hashSet.elements)
}
