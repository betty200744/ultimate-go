package ping_pong

import "fmt"

func PingGen(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
		}
	}()
	return ch
}

// chan<-, send chan
func Ping(chping chan<- string) {
	chping <- "send ping"
	close(chping)
}

//<-chan, receive chan
func Pong(chping <-chan string, chpong chan<- string) {
	chpong <- "receive: " + <-chping + ", then send pong!"
	close(chpong)
}
