package basic_example

import (
	"context"
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/mitchellh/mapstructure"
)

func PrintHostStat(result *api.QueryTableResult) {
	hostStats := make([]*HostStat, 0)
	for result.Next() {
		values := result.Record().Values()
		hostStat := &HostStat{}
		mapstructure.Decode(values, hostStat)
		hostStats = append(hostStats, hostStat)
		fmt.Println(hostStat)
	}
}
func QueryLatestDate() {
	ctx := context.Background()
	client := influxdb2.NewClient("http://localhost:8086", token)
	queryApi := client.QueryAPI("wuyun")
	defer client.Close()
	q := `from(bucket:"stat") 
		|> range(start: -1d) 
		|> filter(fn: (r) => r._measurement == "host_stats") 
		|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> sort(columns: ["_time"], desc: true)
		|> limit(n:4)`
	result, _ := queryApi.Query(ctx, q)
	PrintHostStat(result)
}
