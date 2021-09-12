package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"ultimate-go/awesome-go/gorm/pgtype"
)

type ProductCustom struct {
	gorm.Model
	Code     string                     `json:"code"`
	Price    uint                       `json:"price"`
	JS       pgtype.JSONB               `json:"js"`
	SendRule pgtype.SendRuleStructSlice `json:"send_rule"`
	JsonP    pgtype.Mapping             `gorm:"not null; column:json_p" json:"json_p"`
}

func main() {
	db, err := gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", "postgres", "postgres", "gobyexample"))
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	//db.AutoMigrate(&ProductCustom{})

	// Create
	b, _ := json.Marshal("TEST")
	db.Create(&ProductCustom{Code: "L1213", Price: 1000, JS: b,
		SendRule: pgtype.SendRuleStructSlice{{SendLimit: 1, SendLimitNum: 1}, {SendLimit: 2, SendLimitNum: 2}},
		JsonP:    pgtype.Mapping{T: map[string]interface{}{"a": "a", "b": "b"}},
	})

	// Read
	var product ProductCustom
	db.First(&product)
	fmt.Println(product.SendRule[0])
	fmt.Println(product.JsonP.T)
	// find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	db.First(&product, 1)
}
