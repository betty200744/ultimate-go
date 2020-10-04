package main

import (
	"fmt"

	"github.com/gosidekick/goconfig"
	_ "github.com/gosidekick/goconfig/toml"
)

type Postgresql struct {
	Host string `toml:"host" cfgDefault:"example.com" cfgRequired:"true"`
	Port int    `toml:"port" cfgDefault:"123"`
	Db   string `toml:"db" cfgDefault:"product"`
}

type configTest struct {
	Pg Postgresql `toml:"pg" cfgDefault:"127.0.0.1"`
}

func main() {
	config := configTest{}
	goconfig.File = "config.toml"
	err := goconfig.Parse(&config)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config.Pg)
}
