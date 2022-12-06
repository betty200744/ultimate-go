package ioutil

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	pwd, _ := os.Getwd()
	b, errRead := ioutil.ReadFile(fmt.Sprintf("%s/ioutil_ReadFile.go", pwd))
	fmt.Println(string(b), errRead)
}

func TestWriteFile(t *testing.T) {
	err := ioutil.WriteFile("/Users/betty/project/git_betty200744/ultimate-go/build-in-package/io/my_ioutil/pat2h", []byte(`bbb`), os.ModePerm)
	fmt.Println(err)
}

func Test_TempFile(t *testing.T) {
	os.TempDir()
	file, err := ioutil.TempFile(os.TempDir(), "temp")
	if err != nil {
		panic(err)
	}
	fmt.Println(file.Stat())
}

func Test_ReadAll(t *testing.T) {
	res, _ := http.Get("http://www.baidu.com")
	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Println(string(body))
}
