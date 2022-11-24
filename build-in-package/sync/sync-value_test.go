package sync

import (
	"fmt"
	"sync/atomic"
	"testing"
)

type StringSliceAtomic struct {
	ss atomic.Value
}

func (s *StringSliceAtomic) Load() []string {
	return s.ss.Load().([]string)
}
func (s *StringSliceAtomic) Store(v []string) {
	s.ss.Store(v)
}
func (s *StringSliceAtomic) Add(v string) {
	vv := append(s.ss.Load().([]string), v)
	s.ss.Store(vv)
}
func (s *StringSliceAtomic) First() string {
	vv := s.ss.Load().([]string)
	if len(vv) > 0 {
		return vv[0]
	}
	return ""
}

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
