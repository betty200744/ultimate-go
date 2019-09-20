package main

import (
	"fmt"
	"os"
)

func main() {
	dir := "./build-in-package/os/os_stat.go"
	stat, errStat := os.Stat(dir)
	if errStat != nil {
		fmt.Println(errStat.Error())
	}
	fmt.Println("dir mode: ", stat.Mode())
	fmt.Println("dir mode time:", stat.ModTime())
	fmt.Println("dir isDir: ", stat.IsDir())
}
