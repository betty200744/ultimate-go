package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type concatResponse struct {
	Original string `json:"original"`
	String   string `json:"string"`
	Count    int    `json:"count"`
}

func concat(str string, count int) (ret string) {
	for i := 0; i < count; i++ {
		ret = str + ret
	}
	return ret
}

func concatV2(str string, count int) (ret string) {
	data := bytes.NewBuffer([]byte{})
	for i := 0; i < count; i++ {
		data.WriteString(str)
	}
	return data.String()
}
func ConcatHandler(w http.ResponseWriter, r *http.Request) {
	str := r.FormValue("str")
	count, err := strconv.Atoi(r.FormValue("count"))
	if err != nil {
		count = 10
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp := concatResponse{String: concatV2(str, count), Count: count, Original: str}
	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func GinConcatHandler(c *gin.Context) {
	r := c.Request
	str := r.FormValue("str")
	count, err := strconv.Atoi(r.FormValue("count"))
	if err != nil {
		count = 10
	}
	w := c.Writer
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp := concatResponse{String: concatV2(str, count), Count: count, Original: str}
	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
