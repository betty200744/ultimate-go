package demo

import (
	"fmt"
	"os"
	"testing"
)

func init() {
	os.Setenv("ENVIRON", "localhost")
}
func TestCloudConfigsDemo_GetAll(t *testing.T) {
	cc := New(t)
	fmt.Println(cc.GetAll())

}
func TestCloudConfigsDemo_GetByKey(t *testing.T) {
	cc := New(t)
	fmt.Println(cc.GetAll())
}
