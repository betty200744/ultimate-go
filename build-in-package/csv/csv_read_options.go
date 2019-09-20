package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func main() {
	in := `first_name;last_name;username
"Rob";"Pike";rob
# lines beginning with a # character are ignored
Ken;Thompson;ken
"Robert";"Griesemer";"gri"
`
	ro := csv.NewReader(strings.NewReader(in))
	ro.Comma = ';'
	ro.Comment = '#'
	records, err := ro.ReadAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(records)
}
