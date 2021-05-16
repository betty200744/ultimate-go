package main

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

const (
	DispatchDay = "dispatch_day"
)

func main() {
	ctx := context.Background()
	client := influxdb2.NewClient("http://localhost:8086", "")
	defer client.Close()
	/*
	 Write Data
	*/
	writeApi := client.WriteAPIBlocking("wuyun", "stat")
	p := influxdb2.NewPoint(DispatchDay, map[string]string{"pool_id": "10001", "queue_id": "8"}, map[string]interface{}{"guest_idle": 1, "guest_total": 3}, time.Now())
	// write point immediately
	writeApi.WritePoint(ctx, p)
	// create point using flux style
	pf := influxdb2.NewPointWithMeasurement(DispatchDay).
		AddTag("pool_id", "10001").
		AddTag("queue_id", "7").
		AddField("guest_idle", 1).
		AddField("guest_total", 8).
		SetTime(time.Now())
	writeApi.WritePoint(ctx, pf)
	// write directly line protocol
	line := fmt.Sprintf("dispatch_day,pool_id=10001,queue_id=9 guest_idle=9,guest_total=99")
	writeApi.WriteRecord(ctx, line)
	/*
		Query Data
	*/
	queryApi := client.QueryAPI("wuyun")
	// flux query
	result, err := queryApi.Query(ctx, `from(bucket:"stat") |> range(start: -1h) |> filter(fn: (r) => r._measurement == "dispatch_day") |> limit(n:4)`)
	if err != nil {
		panic(err)
	}
	for result.Next() {
		fmt.Println(result.Record().Values())
	}
}
