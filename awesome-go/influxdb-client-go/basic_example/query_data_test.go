package basic_example

import (
	"context"
	"testing"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func TestQueryBasic(t *testing.T) {
	ctx := context.Background()
	now := time.Now().UnixNano()
	client := influxdb2.NewClient("http://localhost:8086", token)
	queryApi := client.QueryAPI("wuyun")
	defer client.Close()
	WriteHostStatsWithLineProtocol(ctx, "1", "1", "1", 1, 1, 1, now)
	tests := []struct {
		name  string
		query string
	}{
		{
			name: "basic range filter",
			query: `from(bucket:"stat") 
		|> range(start: -1d) 
		|> filter(fn: (r) => r._measurement == "host_stats" and r.bside == "1" and r.pid == "1" ) 
		|> sort(columns: ["_time"], desc: true)
		|> limit(n:4)`,
		},
		{
			name: "basic pivot",
			query: `from(bucket:"stat") 
		|> range(start: -1d) 
		|> filter(fn: (r) => r._measurement == "host_stats" and r.bside == "1" and r.pid == "1" ) 
		|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> sort(columns: ["_time"], desc: true)
		|> limit(n:4)`,
		},
		{
			name: "basic map new table new column with if ",
			query: `from(bucket:"stat") 
		|> range(start: -1d) 
		|> filter(fn: (r) => r._measurement == "host_stats" and r.bside == "1" and r.pid == "1" ) 
		|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> sort(columns: ["_time"], desc: true)
        |> map(fn: (r) => ({pid: r.pid, qid: r.pid, bside: r.bside, used_rate: if r.total == float(v: "0") then float(v: "0") else  r.used / r.total}))
		|> limit(n:4)`,
		},
		{
			name: "basic map with new column",
			query: `from(bucket:"stat") 
		|> range(start: -1d) 
		|> filter(fn: (r) => r._measurement == "host_stats" and r.bside == "1" and r.pid == "1" ) 
		|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> sort(columns: ["_time"], desc: true)
        |> map(fn: (r) => ({r with  used_rate: if r.total == float(v: "0") then float(v: "0") else  r.used / r.total}))
		|> limit(n:4)`,
		},
		{
			name: "basic group, default group by tags",
			query: `from(bucket:"stat") 
		|> range(start: -1d) 
		|> filter(fn: (r) => r._measurement == "host_stats" and r.bside == "1" and r.pid == "1" ) 
		|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> sort(columns: ["_time"], desc: true)
		|> limit(n:4)`,
		},
		{
			name: "basic group, group then sum",
			query: `data = from(bucket:"stat") 
		|> range(start: -1d) 
		|> filter(fn: (r) => r._measurement == "host_stats" and r.bside == "1" and r.pid == "1" ) 
		|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> group(columns: ["pid"], mode: "by")
		
		total = data |> sum(column: "total")
		live = data |> sum(column: "live")
		used = data |> sum(column: "used")
		
		j1 = join(tables: {total: total, live: live}, on: ["pid"])
        join(tables: {j1: j1, used: used}, on: ["pid"])
`,
		},
		{
			name: "basic group, group then mean",
			query: `data = from(bucket:"stat") 
		|> range(start: -1d) 
		|> filter(fn: (r) => r._measurement == "host_stats" and r.bside == "1" and r.pid == "1" ) 
		|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> group(columns: ["pid"], mode: "by")
		
		total = data |> mean(column: "total")
		live = data |> mean(column: "live")
		used = data |> mean(column: "used")
		
		j1 = join(tables: {total: total, live: live}, on: ["pid"])
        join(tables: {j1: j1, used: used}, on: ["pid"])
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := queryApi.Query(ctx, tt.query)
			t.Log(err)
			PrintHostStat(result)
		})
	}
}
