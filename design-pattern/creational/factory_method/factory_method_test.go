package factory_method

import "testing"

func TestGetGun(t *testing.T) {
	ak47, _ := GetGun("ak47")
	PrintDetails(ak47)
	m16, _ := GetGun("m16")
	PrintDetails(m16)
}
