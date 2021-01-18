package user_concurrency

import (
	"testing"
	"time"

	"gobyexample/utils"
)

func TestDetailById(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go DetailById(int64(utils.RandInt(1000, 10000)))
	}
	time.Sleep(time.Minute)
}
