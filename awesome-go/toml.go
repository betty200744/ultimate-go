package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

func main() {
	var blob = `
	key1 = "value1"
	key2 = "value2"
	key3 = "value3"
`

	var cc = `
		key1: 1,
		key2: 2,
		key3: 3,
	`
	type config struct {
		Key1 string
		Key2 string
		Key3 string
	}

	var conf config
	var conf2 config
	md, err := toml.Decode(blob, &conf)
	mcc, err := toml.Decode(cc, &conf2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conf)
	fmt.Println(mcc)
	fmt.Printf("Undecoded keys: %q\n", md.Undecoded())
}
