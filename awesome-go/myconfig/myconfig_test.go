package myconfig

import (
	"fmt"
	"os"
	"testing"
)

type config struct {
	DataSource         string `json:"datasource"`
	Service            string `json:"service"`
	CompanyServiceAddr string `json:"company_service_addr"`
}

func TestLoad(t *testing.T) {
	c := new(config)
	dir, _ := os.Getwd()
	_ = Load(dir+"/localhost.json", c)
	fmt.Println(c)
}
