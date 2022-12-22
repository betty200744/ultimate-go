package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	origin := "http://localhost/"
	url := "ws://localhost:12345/ws"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go ReadMsg(ws)
	go WriteMsg(ws, done)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)
	for {
		select {
		case <-interrupt:
			done <- struct{}{}
			ws.PayloadType = websocket.CloseFrame
			ws.Write([]byte(`write close`))
		}
	}
}

func WriteMsg(ws *websocket.Conn, done chan struct{}) {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case t := <-ticker.C:
			ws.PayloadType = websocket.BinaryFrame
			if _, err := ws.Write([]byte(t.String())); err != nil {
				log.Println("write: ", err)
				return
			}
		case <-done:
			fmt.Println("stop write")
			return
		}
	}
}

func ReadMsg(ws *websocket.Conn) {
	for {
		time.Sleep(time.Second)
		var msg = make([]byte, 512)
		var n int
		n, err := ws.Read(msg)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("Received: %s.\n", msg[:n])
	}
}
