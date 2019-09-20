package main

import (
	"os"
)

func main() {
	dir := "build-in-package/os/export"
	_, errStat := os.Stat(dir)
	if errStat != nil {
		if os.IsNotExist(errStat) {
			os.MkdirAll(dir, os.ModePerm)
		}
	}
}
