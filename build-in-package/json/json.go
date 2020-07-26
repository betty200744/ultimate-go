package main

import (
	"encoding/json"
	"fmt"
	"log"
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

func main() {
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
