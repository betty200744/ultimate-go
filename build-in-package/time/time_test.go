package time

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeParse(t *testing.T) {
	tradeTime := "2020-01-02 15:04:05"
	location, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(time.Parse(layoutMe, tradeTime))
	fmt.Println(time.ParseInLocation(layoutMe, tradeTime, location))
}
