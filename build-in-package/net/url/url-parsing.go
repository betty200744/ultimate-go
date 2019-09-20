package url

import (
	"fmt"
	"net/url"
)

func main() {
	// [scheme:][//[userinfo@]host][/]path[?query][#fragment]
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

}
