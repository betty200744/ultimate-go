package main

import (
	"fmt"
	"time"
)

func foo() {
	fmt.Println("foo...")
}

func main() {
	// time Now,  Sleep, After, Tick
	now := time.Now()
	time.Sleep(time.Microsecond)
	c := <-time.After(1 * time.Second)
	t := time.Tick(time.Second * 2)
	fmt.Println("this is time.After", c)
	for value := range t {
		fmt.Println("this is time.Tick:", value)
	}

	// time.Time Since, Until , After , Before, Add , AddDate, Sub, Year, Month, Day, Hour,Minute, Second, Weekday

	fmt.Println("this is Time.Since:", time.Since(now))
	fmt.Println("this is Time.Until:", time.Until(now)) // since 的反义

	fmt.Println("this is Time.After:", now.After(time.Now()))
	fmt.Println("this is Time.Before:", now.Before(time.Now())) // after 的反义

	fmt.Println("this is Time.Add:", now.Add(time.Second))
	fmt.Println("this is Time.Sub:", now.Sub(time.Now()))
	fmt.Println("this is Time.AddDate:", now.AddDate(0, 0, 1))

	fmt.Println("this is Time.Year():", now.Year())
	fmt.Println("this is Time.Month():", now.Month())
	fmt.Println("this is Time.Day():", now.Day())
	fmt.Println("this is Time.Hour():", now.Hour())
	fmt.Println("this is Time.Minute():", now.Minute())

	// time.Timer NewTimer, AfterFunc , Reset, Stop
	time.AfterFunc(time.Microsecond, foo)

	// time.Duration
	duration := time.Since(now)
	fmt.Println("this is Duration.String:", duration.String())
	fmt.Println("this is Duration.Hours:", duration.Hours())
	fmt.Println("this is Duration.Minutes:", duration.Minutes())
	fmt.Println("this is Duration.Seconds:", duration.Seconds())

}
