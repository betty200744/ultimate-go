package utils

import (
	"fmt"
	"testing"
)

type Review struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

func TestStructToMap(t *testing.T) {
	r := &Review{Id: 1, Name: "n", Owner: "o"}
	res, _ := StructToMap(r)
	fmt.Println(res)
}

func TestGetStringIds(t *testing.T) {

}
