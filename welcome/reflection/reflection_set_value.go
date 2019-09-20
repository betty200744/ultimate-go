package main

import (
	"fmt"
	"reflect"
)

type SkuSimple struct {
	Id             int64
	Name           string
	Price          int64
	RealPrice      int64
	Description    string
	Video          []string
	Imgs           []string
	ImgSnapshot100 string
	Score          int32
	CompanyId      string
	Brand          string
	Poster         string
	Weight         int32
	Barcode        string
	ClassId        int64
	RichText       string
	DiscountPrice  int64
	ActivityTag    []string
	ProductId      string
	CategoryId     int32
	MainCategory   int32
	MiddleCategory int32
	MinorCategory  int32
	SubBrandId     int64
	AboutBrand     string
}

func main() {
	s := struct {
		Id             int64
		Name           string
		Price          int64
		RealPrice      int64
		Description    string
		Video          []string
		Imgs           []string
		ImgSnapshot100 string
		Score          int32
		CompanyId      string
		Brand          string
		Poster         string
		Weight         int32
		Barcode        string
		ClassId        int64
		RichText       string
		DiscountPrice  int64
		ActivityTag    []string
		ProductId      string
		CategoryId     int32
		MainCategory   int32
		MiddleCategory int32
		MinorCategory  int32
		SubBrandId     int64
		AboutBrand     string
	}{Id: 1, Name: "n", Price: 4}

	r := reflect.ValueOf(&s).Elem()

	r.FieldByName("Name").SetString("AAA")
	fmt.Println(s)
}
