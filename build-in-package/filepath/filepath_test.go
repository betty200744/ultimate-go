package filepath

import (
	"fmt"
	"path/filepath"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func Test_FilePathJoin(t *testing.T) {
	res0 := filepath.Join("/home/resource/html", "zfb_page_pay")
	res1 := filepath.Join("/home/resource/html", "/zfb_page_pay")
	res2 := filepath.Join("/home/resource/html", "/zfb_page_pay/")
	assert.Equal(t, "/home/resource/html/zfb_page_pay", res0)
	assert.Equal(t, "/home/resource/html/zfb_page_pay", res1)
	assert.Equal(t, "/home/resource/html/zfb_page_pay", res2)
}

func Test_FilepathDir(t *testing.T) {
	dir := filepath.Dir("/a/bc/abc")
	fmt.Println(dir)
}

func Test_Exists(t *testing.T) {
	res := Exists("/Users/betty/project/git_betty200744/ultimate-go/build-in-package/filepath")
	assert.Equal(t, true, res)
}

func Test_FilePathBase(t *testing.T) {
	uploadHtml := "/home/resource/html"
	res := filepath.Base(uploadHtml)
	assert.Equal(t, "html", res)
}

func Test_FilePathRel(t *testing.T) {
	// relative path
	res, _ := filepath.Rel("/a/b/", "/a/b/c/1.png")
	assert.Equal(t, "c/1.png", res)
}

func TestFilePathAbsPath(t *testing.T) {
	// absolute dir path
	fmt.Println(FilePathAbsPath())
}

func Test_FilePathDir(t *testing.T) {
	res := filepath.Dir("/a/b/c")
	assert.Equal(t, "/a/b", res)
}

func Test_Ext(t *testing.T) {
	// file name extension
	res := filepath.Ext("a.png")
	assert.Equal(t, ".png", res)
}

func TestFilePathAbs(t *testing.T) {
	// absolute file path
	res := FilePathAbs()
	fmt.Println(res)
}
