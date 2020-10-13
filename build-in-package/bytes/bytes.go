package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

func main() {
	type OrderRequestWx struct {
		Appid          string `xml:"appid"`
		MchId          string `xml:"mch_id"`
		NonceStr       string `xml:"nonce_str"`
		OutTradeNo     string `xml:"out_trade_no"`
		SignType       string `xml:"sign_type"`
		Sign           string `xml:"sign"`
		Body           string `xml:"body,omitempty"`
		NotifyUrl      string `xml:"notify_url,omitempty"`
		TradeType      string `xml:"trade_type,omitempty"`
		SpbillCreateIp string `xml:"spbill_create_ip,omitempty"`
		TotalFee       int    `xml:"total_fee,omitempty"`
		// TimeExpire     string `xml:"time_expire,omitempty"`
	}
	param := &OrderRequestWx{
		Appid:          "",
		MchId:          "",
		NonceStr:       "",
		OutTradeNo:     "",
		SignType:       "",
		Sign:           "",
		Body:           "",
		NotifyUrl:      "",
		TradeType:      "",
		SpbillCreateIp: "",
		TotalFee:       0,
	}

	bs, err := xml.Marshal(param)
	if err != nil {
		return
	}
	fmt.Println(string(bs))

	bs = bytes.Replace(bs, []byte("OrderRequestWx"), []byte("xml"), -1)

	fmt.Println(string(bs))

	b := []byte(`ab`)
	bb := []byte{'a', 'b'}
	buffer := bytes.NewBuffer(b)
	fmt.Println(b)
	fmt.Println(bb)
	fmt.Println(buffer)
}
