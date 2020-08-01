package multiplexing

import (
	"fmt"
	"gobyexample/LearnConcurrency/concurrency_visualize/ping_pong"
	"testing"
)

func TestFanIn(t *testing.T) {
	// testcase退出， 则所有goroutine退出
	out := FanIn(ping_pong.PingGen("ping"), ping_pong.PingGen("bing"), ping_pong.PingGen("sing"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-out)
	}
	fmt.Println("aaaaa")
}
