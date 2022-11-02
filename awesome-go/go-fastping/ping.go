package go_fastping

import (
	"fmt"
	"net"
	"time"

	"github.com/tatsushid/go-fastping"
)

func Ping(addr string) {
	p := fastping.NewPinger()
	p.AddIP(addr)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}
	p.OnIdle = func() {
		fmt.Println("finish")
	}
	err := p.Run()
	if err != nil {
		fmt.Println(err)
	}
}
