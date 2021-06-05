package main

import (
	"context"
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/mitchellh/mapstructure"
)

const (
	TsOrg        = "wuyun"
	StatBucket   = "stat"
	HostStatsMea = "host_stats"
	url          = "http://localhost:8086"
	token        = "VmSIO0xg8lpLCBFthbt85ZMQ6_9sjaLRfOg0d7lgQOtv16SzJXzq7fxfVtpt3OjtBywWXQiINwcuY1wHDEwp8w=="
)
const (
	Pid1   = "10000000001"
	Pid2   = "10000000002"
	Qid1   = "1"
	Qid2   = "2"
	Bside1 = "1"
	Bside2 = "2"
)

var (
	client   influxdb2.Client
	writeApi api.WriteAPI
	queryApi api.QueryAPI
)

type HostStats struct {
	// 时间戳 单位秒
	Timestamp int64 `json:"timestamp" mapstructure:"timestamp"`
	// 机房id
	Pid string `json:"pid" mapstructure:"pid"`
	// 业务方，  '1'-小悟云 '2'-领沃
	Bside string `json:"bside" mapstructure:"bside"`
	// 队列id
	Qid string `json:"qid" mapstructure:"qid"`
	// 总设备数，队列折线时无
	Total float64 `json:"total" mapstructure:"total"`
	// 可用设备数，队列折线时无
	Live float64 `json:"live" mapstructure:"live"`
	// 机器使用数
	Used float64 `json:"used" mapstructure:"used"`
}

func QueryTableMetadata(queryApi api.QueryAPI) {
	result, err := queryApi.Query(context.Background(), fmt.Sprintf(`
	data = from(bucket:"%v")
	|> range(start: -1d)
	|> filter(fn: (r) => r._measurement == "%v" and r.pid != "%v" and r.bside != "%v" and r.qid != "%v")
	|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
	|> group(columns: ["qid"], mode:"by")

	total = data |> mean(column: "total")
	live = data |> mean(column: "live")
	used = data |> mean(column: "used")

 	j1 = join(tables: {total: total, live: live}, on: ["qid"])
    join(tables: {j1: j1, used: used}, on: ["qid"])
`, StatBucket, HostStatsMea, "0", "0", "0"))
	if err != nil {
		panic(err)
	}
	for result.Next() {
		if result.TableChanged() {
		}
		s := &HostStats{}
		mapstructure.Decode(result.Record().Values(), s)
		fmt.Println(s)
	}
}

func main() {
	client = influxdb2.NewClient(url, token)
	writeApi = client.WriteAPI(TsOrg, StatBucket)
	queryApi = client.QueryAPI(TsOrg)
	QueryTableMetadata(queryApi)
}
