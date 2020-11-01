package filepath

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func FilePathAbsPath() string {
	_, fileName, _, _ := runtime.Caller(0)
	fmt.Println(fileName)
	return filepath.Dir(fileName)
}
func FilePathAbs() string {
	abs, _ := filepath.Abs("filepath.go")
	return abs
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
