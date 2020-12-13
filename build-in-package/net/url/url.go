package url

import (
	"encoding/base64"
	"fmt"
	"net/url"
)

/*
	url parse
	[scheme:][//[userinfo@]host][/]path[?query][#fragment]
*/
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

func EncodeParam(s string) string {
	return url.QueryEscape(s)
}

func EncodeStringBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
