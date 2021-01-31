package main

import (
	"fmt"
	"net/http"

	"github.com/googollee/go-socket.io"
)

func main() {
	socketio.NewBroadcast()
	server, err := socketio.NewServer(nil)
	if err != nil {
		panic(err)
	}
	server.OnConnect("/", func(conn socketio.Conn) error {
		fmt.Println(conn.URL())
		return nil
	})
	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})
	go server.Serve()
	//http.Handle("/socket.io/", server)
	http.ListenAndServe(":8000", server)
}
