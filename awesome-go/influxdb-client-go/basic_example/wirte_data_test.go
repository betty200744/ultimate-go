package basic_example

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestWriteHostStatsWithLineProtocol(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx   context.Context
		pid   string
		qid   string
		bside string
		total float64
		live  float64
		used  float64
		ctime int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "line protocol 1",
			args: args{ctx: ctx, pid: "1", qid: "1", bside: "1", total: 1, live: 1, used: 1, ctime: time.Now().UnixNano()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteHostStatsWithLineProtocol(tt.args.ctx, tt.args.pid, tt.args.qid, tt.args.bside, tt.args.total, tt.args.live, tt.args.used, tt.args.ctime)
			for i := 0; i < 1000; i++ {
				WriteHostStatsWithLineProtocol(tt.args.ctx, tt.args.pid, tt.args.qid, tt.args.bside, float64(i), float64(i), float64(i), time.Now().Add(time.Minute*time.Duration(-i)).UnixNano())
			}
			fmt.Println("first timestamp is: ", time.Unix(0, tt.args.ctime).Unix())
			QueryLatestDate()
		})
	}
}

func TestWriteHostStatsWithWritePoint(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx   context.Context
		ctime time.Time
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestWriteHostStatsWithWritePoint",
			args: args{
				ctx:   ctx,
				ctime: time.Now(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteHostStatsWithWritePoint(tt.args.ctx, tt.args.ctime)
			fmt.Println("first timestamp is: ", tt.args.ctime.Unix())
			QueryLatestDate()
		})
	}
}

func TestWriteHostStatsWithNewPointWithMeasurement(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx   context.Context
		ctime time.Time
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestWriteHostStatsWithNewPointWithMeasurement",
			args: args{
				ctx:   ctx,
				ctime: time.Now(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteHostStatsWithNewPointWithMeasurement(tt.args.ctx, tt.args.ctime)
			fmt.Println("first timestamp is: ", tt.args.ctime.Unix())
			QueryLatestDate()
		})
	}
}
