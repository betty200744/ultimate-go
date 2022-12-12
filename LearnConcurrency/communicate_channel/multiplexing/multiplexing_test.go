package multiplexing

import (
	"fmt"
	"testing"
	"ultimate-go/LearnConcurrency/communicate_channel/ping_pong"
)

func TestFanIn(t *testing.T) {
	// testcase退出， 则所有goroutine退出
	ping := ping_pong.PingGen("ping")
	bing := ping_pong.PingGen("bing")
	sing := ping_pong.PingGen("sing")
	out := FanIn(ping, bing, sing)
	for i := 0; i < 10; i++ {
		fmt.Println(<-out)
	}
	fmt.Println("aaaaa")
}
func TestFanIn2(t *testing.T) {
	stopPing := make(chan bool)
	ping := ping_pong.PingGenWithQuit("ping", stopPing)
	out := FanIn(ping)
	// manual send stop
	for i := 0; i < 10; i++ {
		fmt.Println(<-out)
	}
	stopPing <- true
	fmt.Println("aaaaa")
}
