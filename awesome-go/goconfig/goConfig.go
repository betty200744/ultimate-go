package main

import (
	"fmt"
	"github.com/crgimenes/goconfig"
)

type Postgresql struct {
	Host string `cfgDefault:"example.com" cfgRequired:"true"`
	Port int    `cfgDefault:"123"`
	Db   string `cfgDefault:"product"`
}

type configTest struct {
	Host string `json:"host" cfgDefault:"127.0.0.1"`
	Pg   Postgresql
}

func main() {
	config := configTest{}
	err := goconfig.Parse(config)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config.Host)
	fmt.Println(config.Pg.Host)
}
