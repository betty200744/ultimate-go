package main

import (
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/t-tiger/gorm-bulk-insert"
)

func BatchInsert(db *gorm.DB, objects []interface{}, excludeColumns ...string) error {
	return gormbulk.BulkInsert(db, objects, 100, excludeColumns...)
}

func TransformToStringArray(items []string) *pq.StringArray {
	stringArray := new(pq.StringArray)
	_ = copier.Copy(&stringArray, &items)
	return stringArray
}

func TransformToInt64Array(items []int64) *pq.Int64Array {
	stringArray := new(pq.Int64Array)
	_ = copier.Copy(&stringArray, &items)
	return stringArray
}

func TransformStringArrayToStringList(pgArray *pq.StringArray) []string {
	var respString []string
	_ = copier.Copy(&respString, &pgArray)
	return respString
}
