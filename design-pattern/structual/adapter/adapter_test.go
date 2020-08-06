package adapter

import "testing"

func TestAdapter_Pattern(t *testing.T) {
	mac := &mac{}
	InsertSquareUsbInComputer(mac)

	win := &windows{}
	winAdapter := &windowsAdapter{windows: win}
	InsertSquareUsbInComputer(winAdapter)
}
