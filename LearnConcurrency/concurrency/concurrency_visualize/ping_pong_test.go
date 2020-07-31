package main

import "testing"

func TestPing_Pong(t *testing.T) {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	Ping(pings)
	Pong(pings, pongs)
}
