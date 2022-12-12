package test

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"sync"
	"ultimate-go/awesome-go/bazel/go-tutorial/service"
)

var (
	router *gin.Engine
	once   sync.Once
)

func init() {
	once.Do(func() {
		router = service.SetUpRouter()
	})
}

func MakeRequest(method, url string, body string) (*httptest.ResponseRecorder, error) {
	b := bytes.NewBufferString(body)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, b)
	header := http.Header{}
	header.Add("Content-Type", " application/json")
	req.Header = header
	router.ServeHTTP(w, req)
	return w, nil
}
