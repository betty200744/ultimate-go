package main

import (
	"testing"
	"time"
)

func TestTickerTimeout(t *testing.T) {
	TickerTimeout(time.Second)
}

func TestTimeAfterTimeout(t *testing.T) {
	TimeAfterTimeout(time.Second)
}
