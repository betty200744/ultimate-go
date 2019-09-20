package main

import (
	"encoding/csv"
	"os"
)

var (
	d *os.File
)

func main() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	n := "/Users/betty/project/whale_go/src/gobyexample/build-in-package/csv/writeAll.csv"
	if _, err := os.Stat(n); os.IsNotExist(err) {
		d, _ = os.Create(n)
	}

	w := csv.NewWriter(d)
	w.WriteAll(records)
}
