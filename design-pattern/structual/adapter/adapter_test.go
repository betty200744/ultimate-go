package adapter

import "testing"

func TestAdapter_Pattern(t *testing.T) {
	mac := &mac{}
	client := &client{}
	client.insertSquareUsbInComputer(mac)

	win := &windows{}
	winAdapter := &windowsAdapter{windows: win}
	client.insertSquareUsbInComputer(winAdapter)
}
