package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

func EchoServer(ws *websocket.Conn) {
	defer ws.Close()
	for {
		req := make([]byte, 512)
		n, err := ws.Read(req)
		if err != nil {
			return
		}
		fmt.Println(ws.Request().RemoteAddr, string(req[:n]))
		ws.Write(req)
	}
}

// This example demonstrates a trivial echo server.
func main() {
	http.Handle("/ws", websocket.Handler(EchoServer))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
