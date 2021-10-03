package main

import (
	"fmt"
	"net/http"
	"time"

	"ultimate-go/utils"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

// conn
// hub
// send chan []byte
const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
	RBufferSize    = 1024
	WBufferSize    = 1024
)

var (
	newline = []byte(`\n`)
	space   = []byte(` `)
)

var upgrader = websocket.Upgrader{
	HandshakeTimeout: 10 * time.Second,
	ReadBufferSize:   RBufferSize,
	WriteBufferSize:  WBufferSize,
}

type User struct {
	id   int64
	conn *websocket.Conn
	hub  *Hub
	send chan []byte
}

func (u *User) readMsg() {
	u.conn.SetReadLimit(maxMessageSize)
	for {
		_, msg, err := u.conn.ReadMessage()
		if err != nil {
			break
		}
		s := fmt.Sprintf("user %d : %v", u.id, string(msg[:]))
		fmt.Println(s)
		u.hub.broadcast <- []byte(s)
	}
}

func (u *User) writeMsg() {
	ticker := time.NewTicker(pingPeriod)
	for {
		select {
		case msg, ok := <-u.send:
			if !ok {
				u.conn.WriteMessage(websocket.CloseMessage, []byte{})
			}
			w, err := u.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(msg)
			n := len(u.send)
			for i := 0; i < n; i++ {
				w.Write(<-u.send)
			}
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			u.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := u.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func WSHandle(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.WithError(err).Error()
	}
	user := &User{
		id:   utils.RandInt64(10000, 99999),
		conn: conn,
		hub:  hub,
		send: make(chan []byte, 512),
	}
	log.Infof("user register userId %d", user.id)
	user.hub.register <- user
	go user.writeMsg()
	go user.readMsg()
}
