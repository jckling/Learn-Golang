package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

func server(ws *websocket.Conn) {
	fmt.Printf("new connection\n") // 建立连接
	buf := make([]byte, 100)
	for {
		if _, err := ws.Read(buf); err != nil {
			fmt.Printf("%s", err.Error())
			break
		}
	}
	fmt.Printf(" => closing connection\n") // 停止连接
	ws.Close()
}

// go run websocket_server.go
func main() {
	http.Handle("/websocket", websocket.Handler(server))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
