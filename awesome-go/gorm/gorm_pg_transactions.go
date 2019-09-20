package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Animal struct {
	Name string `gorm:"not null; column:name" json:"name"`
	Age  int32  `gorm:"not null; column:age" json:"age"`
}

func main() {
	db, err := gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", dbuser, dbpwd, dbname))
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Animal{})

	// Create
	db.Create(&Animal{Name: "L1212", Age: 1000})

	// Read
	var animal Animal
	db.First(&animal, 1)                   // find animal with id 1
	db.First(&animal, "code = ?", "L1212") // find animal with code l1212

	// Update - update animal's price to 2000
	db.Model(&animal).Update("Price", 2000)

	// Delete - delete animal
	db.Delete(&animal)
}
