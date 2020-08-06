package adapter

import "fmt"

// Converts one interface to another so that it matches what the client is expecting
// allows classes with incompatible interfaces to work together by wrapping its own interface around that of an already existing class.
// 封装一个新的类成一个原有的类， 以便兼容
type computer interface {
	insertInSquarePort()
}

func InsertSquareUsbInComputer(com computer) {
	com.insertInSquarePort()
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
