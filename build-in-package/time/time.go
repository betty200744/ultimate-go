package time

import (
	"fmt"
	"time"
)

func TimeProperty() {
	p := fmt.Println
	now := time.Now()
	y, m, d := now.Date()
	p(now)
	p(now.Unix())
	p(now.UnixNano())
	p(now.Year())
	p(y, m, d)
	// property
	fmt.Println("this is Time.Year():", now.Year())
	fmt.Println("this is Time.Month():", now.Month())
	fmt.Println("this is Time.Day():", now.Day())
	fmt.Println("this is Time.Hour():", now.Hour())
	fmt.Println("this is Time.Minute():", now.Minute())
}

func TimeCompare() {
	now := time.Now()
	// compare
	fmt.Println("this is Time.Since:", time.Since(now))
	fmt.Println("this is Time.Until:", time.Until(now)) // since 的反义
	fmt.Println("this is Time.After:", now.After(time.Now()))
	fmt.Println("this is Time.Before:", now.Before(time.Now())) // after 的反义
}

func TimeOperator() {
	now := time.Now()
	since := time.Since(now)
	sub := time.Now().Sub(now)
	fmt.Println(sub)
	// Operator
	fmt.Println("this is Time.Add:", now.Add(time.Second))
	fmt.Println("this is Time.Sub:", now.Sub(time.Now()))
	fmt.Println("this is Time.AddDate:", now.AddDate(0, 0, 1))
	fmt.Println(since, sub)
}

func TimeCron() {
	time.Sleep(time.Microsecond)
	c := <-time.After(1 * time.Second)
	t := time.Tick(time.Second * 2)
	fmt.Println("this is time.After", c)
	for value := range t {
		fmt.Println("this is time.Tick:", value)
	}
	// Cron
	time.AfterFunc(time.Microsecond, func() {
		fmt.Println("foo...")
	})
}
