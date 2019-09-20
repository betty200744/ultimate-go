package main

import (
	"fmt"
	"os"
)

// os.Stat
// file.Create
// file.Chmod
// file.Open //read only
// file.OpenFile // custom perm
// file.Read
// file.Write

func main() {
	path := "./build-in-package/os/export/test"
	_, errStat := os.Stat(path)
	if os.IsNotExist(errStat) {
		os.Create(path)
	}
	if errCh := os.Chmod(path, 0666); errCh != nil {
		fmt.Println(errCh.Error())
	}
	f, errOpen := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if errOpen != nil {
		fmt.Println(errOpen.Error())
	}
	defer f.Close()
	_, errWri := f.Write([]byte(`abc`))
	if errWri != nil {
		fmt.Println(errWri.Error())
	}
}
