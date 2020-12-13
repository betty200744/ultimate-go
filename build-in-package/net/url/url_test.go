package url

import (
	"fmt"
	"log"
	"net/url"
	"testing"
)

func Test_Parse_DB_Url(t *testing.T) {
	postgres := "postgres://user:pass@host.com:5432/path?k=v#f"
	ppostgres, _ := url.Parse(postgres)
	fmt.Println(ppostgres)
	fmt.Println(ppostgres.Host)
	fmt.Println(ppostgres.Path)
	fmt.Println(ppostgres.Query())
	fmt.Println(ppostgres.ForceQuery)
	fmt.Println(ppostgres.Fragment)
	fmt.Println(ppostgres.Scheme)
	fmt.Println(ppostgres.User)
}
func Test_Parse_HttpUrl(t *testing.T) {
	api := "https://api.develop.meetwhale.com/k11/getconfig"
	papi, _ := url.Parse(api)
	fmt.Println(papi)
	fmt.Println(papi.Host)
	fmt.Println(papi.Path)
	fmt.Println(papi.Query())
	fmt.Println(papi.ForceQuery)
	fmt.Println(papi.Fragment)
	fmt.Println(papi.Scheme)
	fmt.Println(papi.User)
}
func Test_ParseQuery(t *testing.T) {
	queryStr := "name=Rajeev%20Singh&phone=%2B9199999999&phone=%2B628888888888"
	params, err := url.ParseQuery(queryStr)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Query Params: ")
	for key, value := range params {
		fmt.Printf("  %v = %v\n", key, value)
	}
}

func Test_BuildStringFromOptions(t *testing.T) {
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

func TestEncodeParam(t *testing.T) {
	s := "https://www.wechatpay.com.cn"
	println(EncodeParam(s))
}

func TestEncodeStringBase64(t *testing.T) {
	s := "https://www.wechatpay.com.cn"
	println(EncodeStringBase64(s))
}
