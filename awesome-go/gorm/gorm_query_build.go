package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	Dbuser = "postgres"
	Dbpwd  = "postgres"
	Dbname = "gobyexample"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func Find(db *gorm.DB, code string, price int, order string) *gorm.DB {
	q := db.Model(&Product{})
	if code != "" {
		q = q.Where("code = ?", code)
	}
	if price != 0 {
		q = q.Where("price = ?", price)
	}
	if order != "" {
		q = q.Order(order)
	}
	return q
}

func main() {
	db, err := gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", Dbuser, Dbpwd, Dbname))
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	// Migrate the schema
	db.AutoMigrate(&Product{})
	// Create
	db.Create(&Product{Code: "L1213", Price: 1000})
	// Find, Builder Pattern
	Find(db, "L1213", 1000, "code DESC")
}
