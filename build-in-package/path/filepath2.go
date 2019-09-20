package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	file, err := ioutil.TempFile(os.TempDir(), "temp")

	if err != nil {
		panic(err)
	}

	fmt.Println("Temp File created!")

	thepath, err := filepath.Abs(filepath.Dir(file.Name()))

	if err != nil {
		panic(err)
	}

	fmt.Println("The file path : ", thepath)

	defer os.Remove(file.Name())

}
