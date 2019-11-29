package main

var (
	x string
)

func f() (string, error) {
	return "f", nil
}
func main() {
	x, _ := f() // x is reassign
	x, _ = f()
	x, y := f()
	x, y = f() // y is not declared
}
