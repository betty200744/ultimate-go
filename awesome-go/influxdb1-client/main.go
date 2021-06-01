package main

import (
	"fmt"

	client "github.com/influxdata/influxdb1-client/v2"
)

func main() {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086"})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	q := client.NewQuery("select * from dispatch_day order by time desc limit 4", "stat", "")
	response, err := c.Query(q)
	fmt.Println(err)
	if err == nil && response.Error() == nil {
		results := response.Results
		for _, r := range results {
			fmt.Println(r)
		}
	}
}
