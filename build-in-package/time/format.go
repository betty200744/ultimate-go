package time

import (
	"fmt"
	"time"
)

// DST, daylight saving time // 夏令时
// CST, Central Standard Time
// MST, Mountain Standard Time
// GMT, Greenwich Mean Time, 就是现在的UTC
// UTC, Coordinated Universal Time
func FormatStandard() {
	now := time.Now()
	fmt.Println("ANSIC: ", now.Format(time.ANSIC))             //  Sun Nov 29 20:09:55 2020
	fmt.Println("UnixDate: ", now.Format(time.UnixDate))       //  Sun Nov 29 20:09:55 CST 2020
	fmt.Println("RFC1123: ", now.Format(time.RFC1123))         //  Sun, 29 Nov 2020 20:09:55 CST
	fmt.Println("RFC3339Nano: ", now.Format(time.RFC3339Nano)) //  2020-11-29T20:09:55.589668+08:00
	fmt.Println("RubyDate: ", now.Format(time.RubyDate))       //  Sun Nov 29 20:09:55 +0800 2020
	fmt.Println("UnixDate: ", now.Format(time.UnixDate))       //  Sun Nov 29 20:09:55 +0800 2020

}
func FormatDemo() {
	currentTime := time.Now()
	fmt.Println(currentTime.Format("2006 年 01 月 02 日"))                                         // 2020 年 11 月 29 日
	fmt.Println("Current Time in String: ", currentTime.String())                               //   2020-11-29 11:15:18.897718 +0800 CST m=+0.000369229
	fmt.Println("MM-DD-YYYY : ", currentTime.Format("01-02-2006"))                              //   11-29-2020
	fmt.Println("YYYY-MM-DD : ", currentTime.Format("2006-01-02"))                              //   2020-11-29
	fmt.Println("YYYY.MM.DD : ", currentTime.Format("2006.01.02 15:04:05"))                     //   2020.11.29 11:15:18
	fmt.Println("YYYY#MM#DD {Special Character} : ", currentTime.Format("2006#01#02"))          //   2020#11#29
	fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))            //   2020-11-29 11:15:18
	fmt.Println("Time with MicroSeconds: ", currentTime.Format("2006-01-02 15:04:05.000000"))   //   2020-11-29 11:15:18.897718
	fmt.Println("Time with NanoSeconds: ", currentTime.Format("2006-01-02 15:04:05.000000000")) //   2020-11-29 11:15:18.897718000
	fmt.Println("ShortNum Month : ", currentTime.Format("2006-1-02"))                           //  2020-11-29
	fmt.Println("LongMonth : ", currentTime.Format("2006-January-02"))                          //   2020-November-29
	fmt.Println("ShortMonth : ", currentTime.Format("2006-Jan-02"))                             //   2020-Nov-29
	fmt.Println("ShortYear : ", currentTime.Format("06-Jan-02"))                                //  20-Nov-29
	fmt.Println("LongWeekDay : ", currentTime.Format("2006-01-02 15:04:05 Monday"))             //   2020-11-29 11:15:18 Sunday
	fmt.Println("ShortWeek Day : ", currentTime.Format("2006-01-02 Mon"))                       //  2020-11-29 Sun
	fmt.Println("ShortDay : ", currentTime.Format("Mon 2006-01-2"))                             //   Sun 2020-11-29
	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5"))           //   2020-11-29 11:15:18
	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5 PM"))        //   2020-11-29 11:15:18 AM
	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5 pm"))        //  2020-11-29 11:15:18 am
}

func FormatBy(layout string) {
	tradeTimeEnd := time.Now().Format(layout)
	fmt.Println(tradeTimeEnd)
}

func FormatGMT() {
	tradeTimeEnd := time.Now().Format("Mon, 02 Jan 2006 15:04:05 GMT")
	fmt.Println(tradeTimeEnd)
}
