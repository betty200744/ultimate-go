package main

import (
	"database/sql/driver"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"strconv"
	"strings"
)

func StringToInt64(a string) int64 {
	int64Val, _ := strconv.ParseInt(a, 10, 64)
	return int64Val
}

type Int8range []int64

func (d *Int8range) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		s := strings.Split(strings.Replace(strings.Replace(string(src), "[", "", -1), ")", "", -1), ",")
		p := make(Int8range, 2)
		p[0] = StringToInt64(s[0])
		p[1] = StringToInt64(s[1])
		*d = p
	case nil:
		return nil
	}
	return nil
}

func (d Int8range) Value() (driver.Value, error) {
	return fmt.Sprintf(
		"[%d, %d)",
		d[0],
		d[1],
	), nil
}

type Booking struct {
	Id             int64     `gorm:"not null; column:id" json:"id"`
	ShowcaseRuleId int64     `gorm:"not null; column:showcase_rule_id" json:"showcase_rule_id"`
	Period         Int8range `gorm:"not null; type:int8range; column:period" json:"period"`
}

type BookingEntity struct {
	Id             int64     `gorm:"not null; column:id" json:"id"`
	ShowcaseRuleId int64     `gorm:"not null; column:showcase_rule_id" json:"showcase_rule_id"`
	Period         Int8range `gorm:"not null; type:int8range; column:period" json:"period"`
}

func main() {
	db, err := gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", "postgres", "postgres", "gobyexample"))
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	//db.AutoMigrate(&Booking{})

	/*
		Create
	*/
	db.Debug()
	db.LogMode(true)
	/*	db.Create(&Booking{
			Id:             1,
			ShowcaseRuleId: 1,
			Int8range:         Int8range{From: 1, To: 3},
		})
		db.Create(&Booking{
			Id:             2,
			ShowcaseRuleId: 1,
			Int8range:         Int8range{From: 4, To: 7},
		})
		db.Create(&Booking{
			Id:             3,
			ShowcaseRuleId: 2,
			Int8range:         Int8range{From: 10, To: 20},
		})*/
	res := new([]Booking)
	// find

	//db.Model(&Booking{}).Find(res)
	db.Model(&Booking{}).Where("lower(period) >= 3::bigint and upper(period) <= 20").Find(res)
	fmt.Println(res)
}
