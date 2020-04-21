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

type Period []int64

func (d *Period) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		s := strings.Split(strings.Replace(strings.Replace(string(src), "[", "", -1), ")", "", -1), ",")
		p := make(Period, 2)
		p[0] = StringToInt64(s[0])
		p[1] = StringToInt64(s[1])
		*d = p
	case nil:
		return nil
	}
	return nil
}

func (d Period) Value() (driver.Value, error) {
	return fmt.Sprintf(
		"[%d, %d)",
		d[0],
		d[1],
	), nil
}

type Booking struct {
	Id             int64  `gorm:"not null; column:id" json:"id"`
	ShowcaseRuleId int64  `gorm:"not null; column:showcase_rule_id" json:"showcase_rule_id"`
	Period         Period `gorm:"not null; type:int8range; column:period" json:"period"`
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
			Period:         Period{1, 3},
		})
		db.Create(&Booking{
			Id:             2,
			ShowcaseRuleId: 1,
			Period:         Period{4, 7},
		})
		db.Create(&Booking{
			Id:             3,
			ShowcaseRuleId: 2,
			Period:         Period{10, 20},
		})*/
	var isCo bool
	var num int32
	res := new([]Booking)
	showcase_rule_ids := make([]int64, 0)
	rulePeriods := []Period{{1, 20}, {4, 7}}

	qSlice := make([]string, 0)
	ruleInterface := make([]interface{}, 0)
	for _, i2 := range rulePeriods {
		ruleInterface = append(ruleInterface, i2)
		qSlice = append(qSlice, "(period && ?)")
	}
	qs := strings.Join(qSlice, " or ")
	fmt.Println(qs)
	db.Model(&Booking{}).Find(res)
	// 存在冲突的条目
	db.Model(&Booking{}).Where("period && ?", rulePeriods[0]).Find(res)
	// 存在冲突的唯一ids
	db.Model(&Booking{}).Where("period && ?", rulePeriods[0]).Pluck("distinct(showcase_rule_id)", &showcase_rule_ids)
	// 存在的条数
	db.Model(&Booking{}).Select("period && ? as collision", rulePeriods[0]).Scan(&isCo)

	db.Model(&Booking{}).Where("lower(period) >= 3::bigint and upper(period) <= 20").Find(res)

	db.Model(&Booking{}).Where("showcase_rule_id = 1").Where(qs, ruleInterface...).Count(&num)

	fmt.Println(res, num, showcase_rule_ids, isCo)
}
