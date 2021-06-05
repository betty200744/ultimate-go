package basic_example

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

const (
	Org         = "wuyun"
	Bucket      = "stat"
	DispatchDay = "dispatch_day"
	HostStats   = "host_stats"
	token       = "VmSIO0xg8lpLCBFthbt85ZMQ6_9sjaLRfOg0d7lgQOtv16SzJXzq7fxfVtpt3OjtBywWXQiINwcuY1wHDEwp8w=="
)

type HostStat struct {
	Pid       string    `json:"pid" mapstructure:"pid"`             // 资源池（虚拟机房）标识
	Qid       string    `json:"qid" mapstructure:"qid"`             // 资源池（虚拟机房）队列标识
	Bside     string    `json:"bside" mapstructure:"bside"`         // 业务方标识(1-小悟云 2-领沃)
	Field     string    `json:"_field" mapstructure:"_field"`       // field
	Value     float64   `json:"_value" mapstructure:"_value"`       // value
	Total     float64   `json:"total" mapstructure:"total"`         // 总数
	Live      float64   `json:"live" mapstructure:"live"`           // 可用数
	Used      float64   `json:"used" mapstructure:"used"`           // 已使用数
	UsedRate  float64   `json:"used_rate" mapstructure:"used_rate"` // 已使用数
	Timestamp time.Time `json:"_time" mapstructure:"_time"`
}

func getClient() influxdb2.Client {
	return influxdb2.NewClient("http://localhost:8086", token)
}

func WriteHostStatsWithWritePoint(ctx context.Context, ctime time.Time) {
	client := getClient()
	writeApi := client.WriteAPIBlocking(Org, Bucket)
	// write point with measurement, tags , fields
	p := influxdb2.NewPoint(HostStats, map[string]string{"pid": "0", "qid": "1", "bside": "1"}, map[string]interface{}{"total": float64(1), "used": float64(3)}, ctime)
	err := writeApi.WritePoint(ctx, p)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()
}
func WriteHostStatsWithNewPointWithMeasurement(ctx context.Context, ctime time.Time) {
	client := getClient()
	writeApi := client.WriteAPIBlocking(Org, Bucket)
	// write point only with measurement, then add tag, add field
	pf := influxdb2.NewPointWithMeasurement(HostStats).
		AddTag("pid", "0").
		AddTag("qid", "0").
		AddTag("bside", "2").
		AddField("total", float64(6)).
		AddField("used", float64(3)).
		SetTime(ctime)
	err := writeApi.WritePoint(ctx, pf)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()
}
func WriteHostStatsWithLineProtocol(ctx context.Context, pid, qid, bside string, total, live, used float64, ctime int64) {
	client := getClient()
	writeApi := client.WriteAPIBlocking(Org, Bucket)
	// write record using line protocol
	line := fmt.Sprintf("%v,pid=%v,qid=%v,bside=%v total=%v,live=%v,used=%v %v", HostStats, pid, qid, bside, total, live, used, ctime)
	err := writeApi.WriteRecord(ctx, line)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()
}
