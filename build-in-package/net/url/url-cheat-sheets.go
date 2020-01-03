package main

import (
	"fmt"
	"net/url"
)

type Options struct {
	URL    string `long:"url" description:"Database connection string"`
	Host   string `long:"host" description:"Server hostname or IP" default:"localhost"`
	Port   int    `long:"port" description:"Server port" default:"5432"`
	User   string `long:"user" description:"Database user"`
	Pass   string `long:"pass" description:"Password for user"`
	DbName string `long:"db" description:"Database name"`
	Ssl    string `long:"ssl" description:"SSL option"`
}

func BuildStringFromOptions(opts Options) (string, error) {
	query := url.Values{}
	if opts.Ssl != "" {
		query.Add("sslmode", opts.Ssl)
	}
	dbUrl := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(opts.User, opts.Pass),
		Host:     fmt.Sprintf("%s:%v", opts.Host, opts.Port),
		Path:     fmt.Sprintf("/%s", opts.DbName),
		RawQuery: query.Encode(),
	}
	return dbUrl.String(), nil
}

func main() {

	/*
		url parse
		[scheme:][//[userinfo@]host][/]path[?query][#fragment]
	*/
	postgres := "postgres://user:pass@host.com:5432/path?k=v#f"
	api := "https://api.develop.meetwhale.com/k11/getconfig"
	ppostgres, _ := url.Parse(postgres)
	papi, _ := url.Parse(api)

	fmt.Println(ppostgres)
	fmt.Println(ppostgres.Host)
	fmt.Println(ppostgres.Path)
	fmt.Println(ppostgres.Query())
	fmt.Println(ppostgres.ForceQuery)
	fmt.Println(ppostgres.Fragment)
	fmt.Println(ppostgres.Scheme)
	fmt.Println(ppostgres.User)

	fmt.Println(papi)
	fmt.Println(papi.Host)
	fmt.Println(papi.Path)
	fmt.Println(papi.Query())
	fmt.Println(papi.ForceQuery)
	fmt.Println(papi.Fragment)
	fmt.Println(papi.Scheme)
	fmt.Println(papi.User)

	/*

		url build
		http : "https://api.develop.meetwhale.com/k11/getconfig"

		postgres : "postgres://user:password@localhost:5432/db?sslmode=require"
	*/

	opts := Options{
		Host:   "localhost",
		Port:   5432,
		User:   "postgres",
		Pass:   "postgres",
		DbName: "product",
		Ssl:    "require",
	}

	str, _ := BuildStringFromOptions(opts)
	fmt.Println(str)

}
