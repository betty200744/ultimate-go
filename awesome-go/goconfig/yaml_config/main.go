package main

import (
	"fmt"

	"github.com/gosidekick/goconfig"
	_ "github.com/gosidekick/goconfig/yaml"
)

type Postgresql struct {
	Host string `yaml:"host" cfgDefault:"example.com" cfgRequired:"true"`
	Port int    `yaml:"port" cfgDefault:"123"`
	Db   string `yaml:"db" cfgDefault:"product"`
}

type configTest struct {
	Pg Postgresql `yaml:"pg" cfgDefault:"127.0.0.1"`
}

func main() {
	config := configTest{}
	goconfig.File = "config.yaml"
	err := goconfig.Parse(&config)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config.Pg)
}
