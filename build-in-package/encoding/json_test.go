package encoding

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Response struct {
	Page  int      `json:"page"`
	Items []string `json:"items"`
}
type FruitBasket struct {
	Name    string    `json:"name"` // default json tag
	Fruit   []string  `json:"fruit"`
	Id      int64     `json:"ref"`     // json tag name ref
	Private string    `json:"private"` // An unexported field is not encoded.
	Created time.Time `json:"created"`
}

func test() {
	// json for  boolean, int, float, string, slice, map, struct
	mapU := make(map[string]string)
	b1, _ := json.Marshal(true)
	i1, _ := json.Marshal(111)
	f1, _ := json.Marshal(3.2)
	s1, _ := json.Marshal("hello")
	sl1, _ := json.Marshal([]string{"a", "b", "c"})
	m1, _ := json.Marshal(map[string]string{"a": "a", "b": "b"})

	stru1, _ := json.Marshal(&Response{
		Page:  1,
		Items: []string{"a", "b", "c"},
	})
	fmt.Println(string(b1))
	fmt.Println(string(i1))
	fmt.Println(string(f1))
	fmt.Println(string(s1))
	fmt.Println(string(sl1))
	fmt.Println(string(m1))
	fmt.Println(string(stru1))
	fmt.Println(json.Unmarshal(m1, &mapU)) // 断言是否解析错误

	// encode
	basket := FruitBasket{
		Name:    "Standard",
		Fruit:   []string{"Apple", "Banana", "Orange"},
		Id:      999,
		Private: "Second-rate",
		Created: time.Now(),
	}

	var jsonData []byte
	jsonData, err := json.Marshal(basket)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))

	// pretty print
	jsonDataPretty, _ := json.MarshalIndent(basket, "", "  ")
	fmt.Println(string(jsonDataPretty))

	//decode
	jsonDataB := []byte(`
{
    "Name": "Standard",
    "Fruit": [
        "Apple",
        "Banana",
        "Orange"
    ],
    "ref": 999,
    "Created": "2018-04-09T23:00:00Z"
}`)

	basketD := new(FruitBasket)
	json.Unmarshal(jsonDataB, basketD)
	fmt.Println(basketD)

	// arbitrary object and arrays
	jsonDataM := []byte(`{
    "Name": "Eve",
    "Age": 6,
    "Parents": [
        "Alice",
        "Bob"
    ]
}`)
	basketM := new(map[string]interface{})
	json.Unmarshal(jsonDataM, basketM)
	fmt.Println(basketM)
}

type Dummy struct {
	Name    string  `json:"name"`
	Number  int64   `json:"number"`
	Pointer *string `json:"pointer"`
}

func test_omitempty() {
	pointer := "yes"

	dummyComplete := Dummy{
		Name:    "Mr Dummy",
		Number:  4,
		Pointer: &pointer,
	}

	data, err := json.Marshal(dummyComplete)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(string(data))

	// ALL of those are considered empty by Go
	dummyNilPointer := Dummy{
		Name:    "",
		Number:  0,
		Pointer: nil,
	}

	dataNil, err := json.Marshal(dummyNilPointer)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(string(dataNil))
}

type ReturnFormat struct {
	ErrNo  int64       `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func Marshal() {
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
