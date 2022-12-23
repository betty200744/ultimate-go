package main

import (
	"bufio"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

func main() {
	u := url.URL{Scheme: "ws", Host: "localhost:12345", Path: "/ws"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.WithError(err).Error()
	}
	defer conn.Close()

	// read message
	go readMsg(conn)

	// read input
	done := make(chan struct{})
	msgCh := make(chan []byte)
	go readStdIn(msgCh, done)

	// write message
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case <-interrupt:
			close(done)
		case msg := <-msgCh:
			conn.WriteMessage(websocket.TextMessage, msg)
		}
	}
}

func readMsg(conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.WithError(err).Error()
		}
		log.Info(string(msg[:]))
	}
}

func readStdIn(msgCh chan []byte, done chan struct{}) {
	for {
		select {
		case <-done:
			return
		default:
			reader := bufio.NewReader(os.Stdin)
			text, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			msgCh <- []byte(text)
		}
	}
}
