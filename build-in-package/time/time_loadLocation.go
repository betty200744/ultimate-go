package main

import (
	"fmt"
	"time"
)

func main() {
	createTime := time.Unix(0, 0)
	shanghai, err := time.LoadLocation("Asia/Shanghai")
	fmt.Println(err)
	createTime.In(shanghai)
	fmt.Println(createTime)
	new_york, err := time.LoadLocation("America/New_York")
	createTime.In(new_york)

	fmt.Println(createTime)

	/*	// China doesn't have daylight saving. It uses a fixed 8 hour offset from UTC.
		secondsEastOfUTC := int((8 * time.Hour).Seconds())
		beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

		// If the system has a timezone database present, it's possible to load a location
		// from that, e.g.:
		//    newYork, err := time.LoadLocation("America/New_York")

		// Creating a time requires a location. Common locations are time.Local and time.UTC.
		timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
		sameTimeInBeijing := time.Date(2009, 1, 1, 20, 0, 0, 0, beijing)

		// Although the UTC clock time is 1200 and the Beijing clock time is 2000, Beijing is
		// 8 hours ahead so the two dates actually represent the same instant.
		//timesAreEqual := timeInUTC.Equal(sameTimeInBeijing)
		//fmt.Println(timesAreEqual)
		fmt.Println(timeInUTC)
		fmt.Println(sameTimeInBeijing)*/

	loc := time.FixedZone("UTC-8", -8*60*60)
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, loc)
	fmt.Println("The time is:", t.Format(time.RFC822))

}
