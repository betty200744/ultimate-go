# goconfig
[![Build Status](https://travis-ci.org/crgimenes/goconfig.svg?branch=master)](https://travis-ci.org/crgimenes/goconfig)
[![Go Report Card](https://goreportcard.com/badge/github.com/crgimenes/goconfig)](https://goreportcard.com/report/github.com/crgimenes/goconfig)
[![Test Coverage](https://api.codeclimate.com/v1/badges/f11c9124505888c4c8db/test_coverage)](https://codeclimate.com/github/crgimenes/goConfig/test_coverage)
[![Maintainability](https://api.codeclimate.com/v1/badges/f11c9124505888c4c8db/maintainability)](https://codeclimate.com/github/crgimenes/goConfig/maintainability)
[![GoDoc](https://godoc.org/github.com/crgimenes/goconfig?status.png)](https://godoc.org/github.com/crgimenes/goconfig)
[![Go project version](https://badge.fury.io/go/github.com%2Fcrgimenes%2Fgoconfig.svg)](https://badge.fury.io/go/github.com%2Fcrgimenes%2Fgoconfig)
[![MIT Licensed](https://img.shields.io/badge/license-MIT-green.svg)](https://tldrlegal.com/license/mit-license)


goconfig uses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file.

## Install

```
go get github.com/crgimenes/goconfig
```

## Example

```go
package main

import "github.com/crgimenes/goconfig"

/*
step 1: Declare your configuration struct,
it may or may not contain substructures.
*/

type mongoDB struct {
	Host string `cfgDefault:"example.com" cfgRequired:"true"`
	Port int    `cfgDefault:"999"`
}

type configTest struct {
	Domain    string
	DebugMode bool `json:"db" cfg:"db" cfgDefault:"false"`
	MongoDB   mongoDB
	IgnoreMe  string `cfg:"-"`
}

func main() {

	// step 2: Instantiate your structure.
	config := configTest{}

	// step 3: Pass the instance pointer to the parser
	err := goconfig.Parse(&config)
	if err != nil {
		println(err)
		return
	}

	/*
	   The parser populated your struct with the data
	   it took from environment variables and command
	   line and now you can use it.
	*/

	println("config.Domain......:", config.Domain)
	println("config.DebugMode...:", config.DebugMode)
	println("config.MongoDB.Host:", config.MongoDB.Host)
	println("config.MongoDB.Port:", config.MongoDB.Port)
}
```

With the example above try environment variables like *$DOMAIN* or *$MONGODB_HOST* and run the example again to see what happens.

You can also try using parameters on the command line, try -h to see the help.

## Contributing

- Fork the repo on GitHub
- Clone the project to your own machine
- Create a *branch* with your modifications `git checkout -b fantastic-feature`.
- Then _commit_ your changes `git commit -m 'Implementation of new fantastic feature'`
- Make a _push_ to your _branch_ `git push origin fantastic-feature`.
- Submit a **Pull Request** so that we can review your changes
