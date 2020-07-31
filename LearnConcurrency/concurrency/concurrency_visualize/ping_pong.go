package main

func Ping(chping chan<- string) {
	chping <- "send ping"
	close(chping)
}

func Pong(chping <-chan string, chpong chan<- string) {
	msg := "receive: " + <-chping + ", then send pong!"
	chpong <- msg
	close(chpong)
}
