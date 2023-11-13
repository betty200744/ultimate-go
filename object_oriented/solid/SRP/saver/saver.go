package saver

import (
	"fmt"
	"os"

	"ultimate-go/object_oriented/solid/SRP"
)

func SaveToFile(fileName string, data SRP.Encoder) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	n, err := f.WriteString(data.Encode())
	if err != nil {
		return err
	}
	fmt.Println(n, "bytes write successfully")
	return nil
}
