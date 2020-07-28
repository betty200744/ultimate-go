package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	type Product struct {
		gorm.Model
		Code  string
		Price int
	}
	db, err := gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", "postgres", "postgres", "gobyexample"))
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	/*
		Create
	*/
	db.Create(&Product{Code: "L1212", Price: int(100)})
	db.Create(&Product{Code: "L1213", Price: int(100)})

	// Read
	var product Product
	db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	/*
		Select
	*/
	products := make([]Product, 0)
	db.Select("code, price").Find(&products)
	fmt.Println(products)

	/*
		Pluck
	*/
	codes := make([]string, 0)
	db.Model(&Product{}).Pluck("code", &codes)
	fmt.Println(codes)

	/*
		Group by code
	*/
	groupResult := new([]struct {
		Code string `json:"code"`
		Sum  int    `json:"sum"`
	})
	db.Table("products").Select("code, sum(price) as sum").Group("code").Scan(groupResult)
	fmt.Println(groupResult)
	// Delete - delete product
	db.Delete(&product)
}
