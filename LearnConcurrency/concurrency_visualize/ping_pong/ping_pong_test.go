package ping_pong

import (
	"fmt"
	"testing"
)

func TestPing_Pong(t *testing.T) {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	// send ping
	Ping(pings)
	Pong(pings, pongs)
	fmt.Println(<-pongs)
}
func TestPingGen(t *testing.T) {
	ch1 := PingGen("ping")
	ch2 := PingGen("bing")
	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
	fmt.Println("PingGen")
}
