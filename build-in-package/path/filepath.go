package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func main() {
	/*	abs,err := filepath.Abs("")
		if err == nil {
			fmt.Println("Absolute:", abs)
		}*/

	_, filename, _, _ := runtime.Caller(1)

	fp := path.Join(path.Dir(filename), "../config/settings.toml")
	fmt.Println(path.Dir(filename))
	fmt.Println(fp)

	path, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}
	dir := filepath.Dir(path)
	fmt.Println(path) // for example /home/user/main
	fmt.Println(dir)  // for example /home/user

	absPath, _ := filepath.Abs("")
	fmt.Println(absPath)

}
