package main

import (
	. "github.com/soekchl/myUtils"
	"github.com/soekchl/websocket"
)

var (
	wsUrl = "ws://localhost:8080/webSocket"
	url   = "http://localhost:8080"
)

func client() {
	ws, err := websocket.Dial(wsUrl, "", url)
	if err != nil {
		Error(err)
		return
	}
	Notice(ws, err)
	ws.Close()
}
