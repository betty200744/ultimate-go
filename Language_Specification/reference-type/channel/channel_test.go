package channel

import (
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	Channel1()
	Channel2()
	Channel3()
	// wait Channel3 goroutine
	time.Sleep(time.Second * 1)
}
