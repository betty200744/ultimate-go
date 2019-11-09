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

	// Create
	db.Create(&Product{Code: "L1212", Price: int(100)})
	db.Create(&Product{Code: "L1213", Price: int(100)})
	type Result struct {
		Code   string  `json:"code"`
		Prices []uint8 `json:"prices"`
	}
	// Read
	var product Product
	r := new([]Result)

	db.Model(&product).Select("code, array_agg(price::int) as prices").Group("code").Scan(&r)
	for _, value := range *r {
		prices := []int{}
		fmt.Println(string(value.Prices))
	}
}
