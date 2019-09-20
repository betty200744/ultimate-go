package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var netTransport = &http.Transport{}
	var client = &http.Client{
		Transport: netTransport,
	}
	resp, _ := client.Get("http://example.com/")
	result, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(result))
}
