package main

import (
	"fmt"

	"github.com/gosidekick/goconfig"
	_ "github.com/gosidekick/goconfig/json"
)

type Postgresql struct {
	Host string `json:"host" cfgDefault:"example.com" cfgRequired:"true"`
	Port int    `json:"port" cfgDefault:"123"`
	Db   string `json:"db" cfgDefault:"product"`
}

type configTest struct {
	Pg Postgresql `json:"pg" cfgDefault:"127.0.0.1"`
}

func main() {
	config := configTest{}
	goconfig.File = "config.json"
	err := goconfig.Parse(&config)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config.Pg)
}
