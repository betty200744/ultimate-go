package main

import (
	"encoding/json"
	"fmt"
)

type ReturnFormat struct {
	ErrNo  int64       `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func main() {
	// data 第一次marshal后存redis
	data := struct {
		Sn      string `json:"sn"`
		Trigger string `json:"trigger"`
	}{Sn: "sn", Trigger: "trigger"}
	dataMarshal, _ := json.Marshal(data)

	// marshal的数据转换成string后， 重新封装再次marshal
	returnUnMarshal := ReturnFormat{}
	returnUnMarshal.ErrNo = 2
	returnUnMarshal.ErrMsg = "hello"
	returnUnMarshal.Data = string(dataMarshal)
	returnMarshal, _ := json.Marshal(returnUnMarshal)

	// 通用函数， 想一次性unmarshal上面分开marshal的数据，但是报错
	result := new(struct {
		ErrNo  int64  `json:"errno"`
		ErrMsg string `json:"errmsg"`
		Data   struct {
			Sn      string `json:"sn"`
			Trigger string `json:"trigger"`
		} `json:"data"`
	})
	err := json.Unmarshal(returnMarshal, result)
	fmt.Println(err)

}
