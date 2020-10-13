package main

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestStringEscape(t *testing.T) {
	StringEscape()
}

func TestSnakeCaseToCamelCase(t *testing.T) {
	snakeCase := "img_snapshot_100"
	res := SnakeCaseToCamelCase(snakeCase)
	assert.Equal(t, "ImgSnapshot100", res)
}
func TestCamelCaseToUnderscore(t *testing.T) {
	res := CamelCaseToUnderscore("createTime")
	assert.Equal(t, "create_time", res)
}

func TestStringsReader(t *testing.T) {
	StringsReader()
}
func TestStringTrim(t *testing.T) {
	StringTrim()
}
func TestTrimSuffix(t *testing.T) {

}
