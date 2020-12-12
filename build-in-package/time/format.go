package time

import (
	"fmt"
	"time"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
	layoutGMT = "Mon, 02 Jan 2006 15:04:05 GMT"
	layoutMe  = "2006-01-02 15:04:05"
)

func FormatStandard() {
	now := time.Now()
	fmt.Println("ANSIC: ", now.Format(time.ANSIC))             //  Sun Nov 29 20:09:55 2020
	fmt.Println("UnixDate: ", now.Format(time.UnixDate))       //  Sun Nov 29 20:09:55 CST 2020
	fmt.Println("RFC1123: ", now.Format(time.RFC1123))         //  Sun, 29 Nov 2020 20:09:55 CST
	fmt.Println("RFC3339Nano: ", now.Format(time.RFC3339Nano)) //  2020-11-29T20:09:55.589668+08:00
	fmt.Println("RubyDate: ", now.Format(time.RubyDate))       //  Sun Nov 29 20:09:55 +0800 2020
	fmt.Println("UnixDate: ", now.Format(time.UnixDate))       //  Sun Nov 29 20:09:55 +0800 2020

}

func FormatBy(layout string) {
	tradeTimeEnd := time.Now().Format(layout)
	fmt.Println(tradeTimeEnd)
}

func FormatGMT() {
	tradeTimeEnd := time.Now().Format("Mon, 02 Jan 2006 15:04:05 GMT")
	fmt.Println(tradeTimeEnd)
}
