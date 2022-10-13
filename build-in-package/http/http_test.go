package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func Test_Get(t *testing.T) {
	res, _ := Get("https://jsonplaceholder.typicode.com/users/1")
	fmt.Println(string(res[:]))
}
func Test_Get_2(t *testing.T) {
	resp, _ := http.Get("http://example.com/")
	result, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(result))
}
func TestGetImage(t *testing.T) {
	res, _ := GetImage("https://bit.ly/2IRnmVm")
	fmt.Println(res)
}

func TestGetWithHeader(t *testing.T) {
	url := "http://example.com/"
	headers := map[string]string{"Authorization": "Basic YWRtaW46QnV6aG9uZ3lhbzEyMw=="}
	res, _ := GetWithHeader(url, headers)
	fmt.Println(string(res[:]))
}
func TestPostJson(t *testing.T) {
	url := "http://dummy.restapiexample.com/api/v1/create"
	payload := `		{
			"name":"test",
			"salary":"123",
			"age":"23"
		}`
	res, _ := PostJson(url, []byte(payload))
	fmt.Println(string(res[:]))
}
func TestDoPost(t *testing.T) {
	url := "http://dummy.restapiexample.com/api/v1/create"
	payload := `		{
			"name":"test",
			"salary":"123",
			"age":"23"
		}`
	res, _ := DoPost(TypeJson, url, nil, []byte(payload))
	fmt.Println(string(res[:]))
}
func Test_FileServer(t *testing.T) {
	err := http.ListenAndServe("localhost:3005", http.FileServer(http.Dir("./")))
	if err != nil {
		log.Fatal(err)
	}
}
func Test_NewServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/heart", HeartHandle)
	mux.HandleFunc("/register", RegisterHandle)
	mux.HandleFunc("/heartbeat", HeartbeatHandle)
	mux.HandleFunc("/statistics", StatisticsHandle)
	mux.HandleFunc("/status", StatusHandle)
}

func TestPostFile(t *testing.T) {
	res, _ := PostFile("https://google.com/upload", nil, "file", "/Users/betty/project/git_betty200744/ultimate-go/img/ciphers.png")
	fmt.Println(string(res[:]))
}
