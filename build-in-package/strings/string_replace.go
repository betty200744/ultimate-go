package main

import (
	"fmt"
	"strings"
)

func main() {
	//order := "UPDATE_TIME__DESC"
	//order := "ACTIVITIES__DETAIL_TYPE__DESC"
	order := "BB__ACTIVITIES__DETAIL_TYPE__DESC"
	nestedPath := ""
	sortField := "update_time"
	ascending := false

	items := strings.Split(order, "__")
	if len(items) > 2 {
		nestedPath = strings.ToLower(strings.Join(items[:len(items)-2], "."))
		sortField = strings.ToLower(nestedPath + "." + items[len(items)-2])
	} else {
		sortField = strings.ToLower(items[len(items)-2])
	}

	switch items[len(items)-1] {
	case "ace":
		ascending = true
	case "desc":
		ascending = false
	}
	fmt.Println("nested path:", nestedPath)
	fmt.Println("sortField:", sortField)
	fmt.Println("sort order:", ascending)

	s :=
		strings.ReplaceAll(s, "\"\"", "|")
}
