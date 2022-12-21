package encoding

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func TestCSVRead(t *testing.T) {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	r := csv.NewReader(strings.NewReader(in))
	for {

		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}
}

func TestCSVWrite(t *testing.T) {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}
	file, err := os.Create("./a.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	w := csv.NewWriter(file)
	for _, record := range records {
		if e := w.Write(record); e != nil {
			log.Fatal(e)
		}
	}
	w.Flush()
}
