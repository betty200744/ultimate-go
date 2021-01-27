package atomic

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func Test_Value(t *testing.T) {
	var AtomicStringSlice atomic.Value
	AtomicStringSlice.Store([]string{"a", "b", "c"})
	v := AtomicStringSlice.Load()
	for _, i2 := range v.([]string) {
		fmt.Println(i2)
	}
}

func TestStringSliceAtomic_Load(t *testing.T) {
	s := &StringSliceAtomic{}
	s.Store([]string{"a", "b", "c"})
	vv := s.Load()
	for _, v := range vv {
		fmt.Println(v)
	}
	fmt.Println("first: ", s.First())
	s.Add("d")
	vv = s.Load()
	for _, v := range vv {
		fmt.Println(v)
	}
}
