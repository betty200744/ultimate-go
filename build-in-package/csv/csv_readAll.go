package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func main() {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	ra := csv.NewReader(strings.NewReader(in))
	records, err := ra.ReadAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(records)
}
