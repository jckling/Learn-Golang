package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"time"
)

// go run websocket_client.go
func main() {
	ws, err := websocket.Dial("ws://localhost:8080/websocket", "",
		"http://localhost/")
	if err != nil {
		panic("Dial: " + err.Error())
	}
	go readFromServer(ws)
	time.Sleep(5e9)
	ws.Close()
}

func readFromServer(ws *websocket.Conn) {
	buf := make([]byte, 1000)
	for {
		if _, err := ws.Read(buf); err != nil {
			fmt.Printf("%s\n", err.Error())
			break
		}
	}
}
