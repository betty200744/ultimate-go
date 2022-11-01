package reflection_ValueOf

import (
	"fmt"
	"testing"
)

func TestRint(t *testing.T) {
	Rint()
}
func TestRstring(t *testing.T) {
	Rstring()
}
func TestRarray(t *testing.T) {
	Rarray()
}
func TestRmap(t *testing.T) {
	Rmap()
}
func TestRfun(t *testing.T) {
	Rfun()
}
func TestRstruct(t *testing.T) {
	Rstruct()
}

func TestStructFieldMethod_ReceiveMethod1(t *testing.T) {
	s1 := Student{1, 1, "Jack", 22}
	s2 := Student{2, 2, "betty", 22}
	ss := []*Student{&s1, &s2}
	ids := make([]int64, len(ss))
	ims := make([]int32, len(ss))
	for i, s := range ss {
		ids[i] = getValueByField("Id", s).(int64)
		ims[i] = getValueByField("ImagId", s).(int32)
	}
	fmt.Println(ids)
	fmt.Println(ims)
}
