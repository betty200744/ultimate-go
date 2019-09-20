package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Id       int64
	Name     string
	Channels []string
}

func main() {
	input1 := map[string]interface{}{
		"id":       1,
		"name":     "betty",
		"channels": []string{"c1", "c2"},
	}
	input2 := map[string]interface{}{
		"id":       1,
		"name":     "betty",
		"channels": []string{"c1", "c2"},
	}
	inputs := []map[string]interface{}{input1, input2}
	output := &[]Person{}
	mapstructure.Decode(inputs, output)
	fmt.Println(output)
}
