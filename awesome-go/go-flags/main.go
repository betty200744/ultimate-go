package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

type Options struct {
	Version bool   `short:"v" long:"version" description:"Print version"`
	Debug   bool   `short:"d" long:"debug" description:"Enable debugging mode"`
	URL     string `long:"url" description:"Database connection string"`
	Host    string `long:"host" description:"Server hostname or IP" default:"localhost"`
	Port    int    `long:"port" description:"Server port" default:"5432"`
}

func main() {
	/*
		go run ./main.go --version true
		{true false  localhost 5432}

	*/
	var opts = Options{}
	args := os.Args
	_, err := flags.ParseArgs(&opts, args)
	if err != nil {
		fmt.Println(err)
	}
	//
	fmt.Println(opts)
}
