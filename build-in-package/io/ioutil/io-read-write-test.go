package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	pwd, _ := os.Getwd()
	b, errRead := ioutil.ReadFile(fmt.Sprintf("%s/build-in-package/io-read-write-test.go", pwd))
	fmt.Println(string(b), errRead)
	res, _ := http.Get("http://www.baidu.com")
	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Println(string(body))
}
