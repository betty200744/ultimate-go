package user_concurrency

import (
	"fmt"
	"sync"
	"time"
)

const (
	QueryConcurrency = 10
)

var (
	QueryAccess chan struct{}
	once        sync.Once
)

func init() {
	once.Do(func() {
		QueryAccess = make(chan struct{}, QueryConcurrency)
		for i := 0; i < QueryConcurrency; i++ {
			QueryAccess <- struct{}{}
		}
	})
}

func DetailById(id int64) int64 {
	// 访问控制
	select {
	case a := <-QueryAccess:
		defer func() {
			QueryAccess <- a
		}()
	default:
		panic("TooFrequencyErr")
	}
	fmt.Println(id)
	time.Sleep(time.Millisecond)
	return id
}
