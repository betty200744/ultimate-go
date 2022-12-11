package main

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/sync/singleflight"
)

func Test_singleflight(t *testing.T) {
	var sf singleflight.Group
	ll := []string{"a", "b", "c", "a", "b", "c"}
	for _, s := range ll {
		go func(s string) {
			v, _, shared := sf.Do(s, func() (interface{}, error) {
				HelloWorld(s)
				return s, nil
			})
			t.Logf(fmt.Sprintf("single flight, s: %s  v: %v , shared:%v ", s, v, shared))
		}(s)
	}
	time.Sleep(time.Second * 3)
}
