package ioutil

import (
	"io/ioutil"
	"os"
)

func RemoveEmptyDir(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}
	if len(files) != 0 {
		return
	}
	err = os.Remove(path)
	if err != nil {
		return
	}
}
