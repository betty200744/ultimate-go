package time

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeOutByTimer(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
func Test_Timer_Retry(t *testing.T) {
	ttl := time.Second * 3
	deadline := time.Now().Add(time.Second * 7)
	var maxRetry int64 = 2
	var timer *time.Timer
	for {
		// timeout
		if !time.Now().Before(deadline) {
			fmt.Println("total duration timeout")
			return
		}

		// first
		if timer == nil {
			timer = time.NewTimer(time.Duration(int64(ttl) / maxRetry))
		} else {
			// retry
			timer.Reset(ttl)
		}
		// max retry time
		if maxRetry > 1 {
			maxRetry--
		} else {
			fmt.Println("max retry end")
			return
		}
		select {
		case <-timer.C:
			fmt.Println(time.Now())
		}
	}
}

func TestTimeoutByRetry(t *testing.T) {
	ttl := time.Second * 3
	deadline := time.Now().Add(time.Second * 7)
	maxRetry := int64(10)
	type args struct {
		ttl      time.Duration
		deadline time.Time
		maxRetry int64
		fn       func(ttl time.Duration) bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TimeoutByRetry success",
			args: args{
				ttl:      ttl,
				deadline: deadline,
				maxRetry: maxRetry,
				fn: func(ttl time.Duration) bool {
					return true
				},
			},
		},
		{
			name: "TimeoutByRetry timeout",
			args: args{
				ttl:      ttl,
				deadline: deadline,
				maxRetry: maxRetry,
				fn: func(ttl time.Duration) bool {
					return false
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TimeoutByRetry(tt.args.ttl, tt.args.deadline, tt.args.maxRetry, tt.args.fn)
		})
	}
}
