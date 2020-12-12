package time

import (
	"fmt"
	"testing"
	"time"
)

func TestLocation(t *testing.T) {
	currentTime := time.Now()
	fmt.Println(currentTime)                            // default, current local time
	fmt.Println(currentTime.Local())                    // local time, Central Standard Time
	fmt.Println(currentTime.UTC())                      // utc or cst time
	fmt.Println(currentTime.Location())                 // the time zone information
	fmt.Println(currentTime.Format("2006 年 01 月 02 日")) // 2020 年 11 月 29 日
}

func TestLoadLocation(t *testing.T) {
	tradeTime := "2020-01-02 15:04:05"
	location, _ := time.LoadLocation("Asia/Shanghai")
	time.Parse(layoutMe, tradeTime)
	time.ParseInLocation(layoutMe, tradeTime, location)
}
