package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

/*
1. only exported Fields can be json output
2. use tag name when has struct tag name
3. point nil
*/

type FruitBasket struct {
	Name    string // default json tag
	Fruit   []string
	Id      int64  `json:"ref"` // json tag name ref
	private string // An unexported field is not encoded.
	Created time.Time
}

func main() {

	// encode
	basket := FruitBasket{
		Name:    "Standard",
		Fruit:   []string{"Apple", "Banana", "Orange"},
		Id:      999,
		private: "Second-rate",
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
