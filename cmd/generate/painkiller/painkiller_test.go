package painkiller

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {

	m.Run()
}
func TestPill_String(t *testing.T) {
	fmt.Printf("%d \n", Paracetamol)
	fmt.Println(Paracetamol.String())
}
