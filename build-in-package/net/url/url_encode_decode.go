package main

import (
	"encoding/base64"
	"net/url"
)

func main() {
	s := "https://www.wechatpay.com.cn"
	println(EncodeParam(s))

	println(EncodeStringBase64(s))

}
func EncodeParam(s string) string {
	return url.QueryEscape(s)
}

func EncodeStringBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
