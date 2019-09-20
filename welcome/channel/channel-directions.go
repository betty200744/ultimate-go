package main

import "fmt"

// chan , send and receive chan
// chan<-, send chan
// <-chan, receive chan

func ping(chping chan<- string, msg string) {
	fmt.Println(msg)
	chping <- msg
}

func pong(chpong chan<- string, chping <-chan string) {
	msg := "200ok, " + <-chping
	chpong <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "this is ping")
	pong(pongs, pings)
	fmt.Println(<-pongs)
}
