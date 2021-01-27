package atomic

import (
	"sync/atomic"
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
