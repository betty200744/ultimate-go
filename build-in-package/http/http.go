package http

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const (
	TypeXml  = 0
	TypeForm = 1
	TypeJson = 2
)

var (
	once   sync.Once
	client *http.Client
)

func newHTTPClient() *http.Client {
	once.Do(func() {
		client = &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost: 10,
				MaxConnsPerHost:     10,
			},
		}
	})
	return client
}
func Get(url string) ([]byte, error) {
	res, errRes := http.Get(url)
	if errRes != nil {
		return nil, errRes
	}
	data, errRead := ioutil.ReadAll(res.Body)
	if errRead != nil {
		return nil, errRead
	}
	defer res.Body.Close()
	return data, nil
}
func GetImage(url string) (string, error) {
	res, errRes := http.Get(url)
	if errRes != nil {
		log.Fatal(errRes)
	}
	data, errRead := ioutil.ReadAll(res.Body)
	if errRead != nil {
		return "", errRead
	}
	defer res.Body.Close()
	filename := "keyboard.jpeg"
	if res.Header.Get("Content-Type") == "image/jpeg" {
		ioutil.WriteFile(filename, data, os.ModePerm)
	}
	return filename, nil
}
func GetWithHeader(fullQuery string, headerInput map[string]string) (response []byte, err error) {
	req, _ := http.NewRequest("GET", fullQuery, strings.NewReader(""))
	req.Header.Set("Content-Type", "application/json")
	for key, val := range headerInput {
		req.Header.Set(key, val)
	}
	client := newHTTPClient()
	res, errRes := client.Do(req)
	if errRes != nil {
		return nil, errRes
	}
	defer res.Body.Close()
	body, errRead := ioutil.ReadAll(res.Body)
	if errRead != nil {
		return response, errRead
	}
	return body, nil
}
func PostJson(url string, payload []byte) ([]byte, error) {
	client = newHTTPClient()
	res, err := client.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
func DoPost(formType int, url string, header map[string]string, body []byte) (bs []byte, err error) {
	var req *http.Request
	req, err = http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return
	}
	if header == nil {
		if formType == TypeXml {
			req.Header.Set("Accept", "application/xml")
			req.Header.Set("Content-Type", "application/xml;charset=utf-8")
		} else if formType == TypeForm {
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
		} else if formType == TypeJson {
			req.Header.Set("Content-Type", "application/json;charset=utf-8")
		}
	} else {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	var resp *http.Response
	c := http.Client{}
	resp, err = c.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = fmt.Errorf("StatusCode=%d", resp.StatusCode)
		return
	}
	bs, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
func PostFile(uri string, params map[string]string, paramName, path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", uri, body)
	client := newHTTPClient()
	res, resErr := client.Do(req)
	if resErr != nil {
		return nil, err
	}
	data, errRead := ioutil.ReadAll(res.Body)
	if errRead != nil {
		return nil, errRead
	}
	return data, nil
}
func HeartHandle(write http.ResponseWriter, request *http.Request) {
	_, _ = write.Write([]byte(request.URL.Path))
	return
}
func RegisterHandle(write http.ResponseWriter, request *http.Request) {
	_, _ = write.Write([]byte(request.URL.Path))
	return
}
func HeartbeatHandle(write http.ResponseWriter, request *http.Request) {
	_, _ = write.Write([]byte(request.URL.Path))
	return
}
func StatisticsHandle(write http.ResponseWriter, request *http.Request) {
	_, _ = write.Write([]byte(request.URL.Path))
	return
}
func StatusHandle(write http.ResponseWriter, request *http.Request) {
	_, _ = write.Write([]byte(request.URL.Path))
	return
}
