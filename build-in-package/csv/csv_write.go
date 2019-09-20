package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

var (
	f *os.File
)

func main() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}
	n := "/Users/betty/project/whale_go/src/gobyexample/build-in-package/csv/write.csv"

	if _, err := os.Stat(n); os.IsNotExist(err) {
		f, _ = os.Create(n)
	}
	w := csv.NewWriter(f)
	for _, value := range records {
		errW := w.Write(value)
		if errW != nil {
			fmt.Println(errW)
		}
	}
	w.Flush()
}
