package main

import (
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	gnow "github.com/jinzhu/now"
)

const (
	TsOrg      = "wuyun"
	StatBucket = "stat"
	url        = "http://localhost:8086"
	token      = "VmSIO0xg8lpLCBFthbt85ZMQ6_9sjaLRfOg0d7lgQOtv16SzJXzq7fxfVtpt3OjtBywWXQiINwcuY1wHDEwp8w=="
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

	now           = time.Now()
	oneHourBefore = time.Now().Add(-time.Hour)
	oneDayBefore  = time.Now().AddDate(0, 0, -1)
	twoDayBefore  = time.Now().AddDate(0, 0, -2)
	beginOfDay    = gnow.With(now).BeginningOfDay().Add(22 * time.Hour)
	beginOfMon    = gnow.With(now).BeginningOfMonth()
	endOfDay      = gnow.With(now).EndOfDay()
	endOfMon      = gnow.With(now).EndOfMonth()
	min5          = beginOfDay.Add(time.Minute * 5).UnixNano()
	min10         = beginOfDay.Add(time.Minute * 10).UnixNano()
	min15         = beginOfDay.Add(time.Minute * 15).UnixNano()
	min20         = beginOfDay.Add(time.Minute * 20).UnixNano()
	day1          = beginOfMon.AddDate(0, 0, 21).UnixNano()
	day2          = beginOfMon.AddDate(0, 0, 22).UnixNano()
	day3          = beginOfMon.AddDate(0, 0, 23).UnixNano()
	day4          = beginOfMon.AddDate(0, 0, 24).UnixNano()
)

func HostStat() {
	client = influxdb2.NewClient(url, token)
	queryApi = client.QueryAPI(TsOrg)
	writeApi = client.WriteAPI(TsOrg, StatBucket)
	bsides := []string{Bside1, Bside2}
	pids := []string{Pid1, Pid2}
	qids := []string{Qid1, Qid2}
	mins := []int64{min5, min10, min15, min20}
	linemin := "host_stats,pid=%v,bside=%v,qid=%v,instance_id=1 total=%v,live=%v,used=%v %d"
	for _, min := range mins {
		// 1, all
		writeApi.WriteRecord(fmt.Sprintf(linemin, 0, 0, 0, 16, 8, 8, min))
		for _, bside := range bsides {
			// 2, bside dimension
			writeApi.WriteRecord(fmt.Sprintf(linemin, 0, bside, 0, 8, 4, 4, min))
			for _, pid := range pids {
				// 5, pid & bside & qid
				for _, qid := range qids {
					writeApi.WriteRecord(fmt.Sprintf(linemin, pid, bside, qid, 2, 1, 1, min))
				}
			}
		}
		for _, pid := range pids {
			// 3, pid & bside 0 dimension
			writeApi.WriteRecord(fmt.Sprintf(linemin, pid, 0, 0, 8, 4, 4, min))
		}
		for _, pid := range pids {
			// 4, pid  & qid & bside 0
			for _, qid := range qids {
				writeApi.WriteRecord(fmt.Sprintf(linemin, pid, 0, qid, 4, 2, 2, min))
			}
		}

	}
	writeApi.Flush()
}
func HostStatDay() {
	bsides := []string{Bside1, Bside2}
	pids := []string{Pid1, Pid2}
	qids := []string{Qid1, Qid2}
	days := []int64{day1, day2, day3, day4}
	lineday := "host_stats_day,pid=%v,bside=%v,qid=%v total=%v,live=%v,used=%v %d"
	for _, day := range days {
		// 1, all
		writeApi.WriteRecord(fmt.Sprintf(lineday, 0, 0, 0, 16, 8, 8, day))
		for _, bside := range bsides {
			// 2, bside dimension
			writeApi.WriteRecord(fmt.Sprintf(lineday, 0, bside, 0, 8, 4, 4, day))
			for _, pid := range pids {
				// 5, pid & bside & qid
				for _, qid := range qids {
					writeApi.WriteRecord(fmt.Sprintf(lineday, pid, bside, qid, 2, 1, 1, day))
				}
			}
		}

		for _, pid := range pids {
			// 3, pid & bside 0 & qid = 0 dimension
			writeApi.WriteRecord(fmt.Sprintf(lineday, pid, 0, 0, 8, 4, 4, day))
		}
		for _, pid := range pids {
			// 4, pid  & qid & bside 0
			for _, qid := range qids {
				writeApi.WriteRecord(fmt.Sprintf(lineday, pid, 0, qid, 4, 2, 2, day))
			}
		}

	}
	writeApi.Flush()
}

func main() {
	HostStat()
	HostStatDay()
}
