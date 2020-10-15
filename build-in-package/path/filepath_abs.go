package main

import (
	"fmt"
	"path"
	"runtime"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println(path.Join(path.Dir(filename)))
}
