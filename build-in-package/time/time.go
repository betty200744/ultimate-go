package time

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	yesterday := time.Date(2019, 2, 1, 12, 30, 0, 0, time.UTC)
	y, m, d := now.Date()
	p(now)
	p(now.Unix())
	p(now.UnixNano())
	p(yesterday)
	p(now.Year())
	p(y, m, d)
	p(yesterday.Before(now))
	p(now.Sub(yesterday))
	p(time.Since(now), time.Now().Sub(now))

	p(yesterday.Add(-time.Hour * 2))

}
