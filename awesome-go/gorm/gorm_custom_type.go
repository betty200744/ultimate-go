package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// ######################

type SendRuleStruct struct {
	SendLimit    int32 `json:"send_limit"`
	SendLimitNum int64 `json:"send_limit_num"`
}
type SendRuleStructSlice []*SendRuleStruct

func (p SendRuleStructSlice) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

func (p *SendRuleStructSlice) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	i := new([]*SendRuleStruct)
	if err := json.Unmarshal(source, i); err != nil {
		return err
	}

	*p = *i
	return nil
}

// #######################

func (p SendRuleStruct) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *SendRuleStruct) Scan(src interface{}) (err error) {
	source, ok := src.([]byte)
	if !ok {
		err = errors.New("Type assertion .([]byte) failed.")
		return
	}
	err = json.Unmarshal(source, p)
	return
}

// ##############
type JSONB []byte

func (j JSONB) Value() (driver.Value, error) {
	return string(j), nil
}
func (j *JSONB) Scan(value interface{}) error {
	s, ok := value.([]byte)
	if !ok {
		errors.New("Scan source was not string")
	}
	*j = append((*j)[0:0], s...)

	return nil
}

// ##############
type JsonP struct {
	T map[string]interface{}
}

func (p JsonP) Value() (driver.Value, error) {
	j, _ := json.Marshal(p.T)
	return j, nil
}
func (p *JsonP) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	var i map[string]interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	p.T = i
	return nil
}

type ProductCustom struct {
	gorm.Model
	Code     string              `json:"code"`
	Price    uint                `json:"price"`
	JS       JSONB               `json:"js"`
	SendRule SendRuleStructSlice `json:"send_rule"`
	JsonP    JsonP               `json:"json_p"`
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
		SendRule: SendRuleStructSlice{{SendLimit: 1, SendLimitNum: 1}, {SendLimit: 2, SendLimitNum: 2}},
		JsonP:    JsonP{T: map[string]interface{}{"a": "a", "b": "b"}},
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
