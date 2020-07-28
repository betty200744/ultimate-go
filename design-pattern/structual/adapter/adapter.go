package adapter

import "fmt"

type computer interface {
	insertInSquarePort()
}

type mac struct {
}

func (m *mac) insertInSquarePort() {
	fmt.Println("Insert Square port into mac")
}

type windows struct {
}

func (w *windows) insertInCirclePort() {
	fmt.Println("Insert Circle port into windows")
}

type windowsAdapter struct {
	*windows
}

func (w *windowsAdapter) insertInSquarePort() {
	w.windows.insertInCirclePort()
}

type client struct {
}

func (c *client) insertSquareUsbInComputer(com computer) {
	com.insertInSquarePort()
}
