package main

import (
	"./rpc_objects"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

// go run rpc_server.go
func main() {
	calc := new(rpc_objects.Args)
	rpc.Register(calc) // 注册对象
	// rpc.RegisterName("Calculator", calc) // 按名称注册
	rpc.HandleHTTP()

	listener, e := net.Listen("tcp", "localhost:8080") // 监听
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}

	go http.Serve(listener, nil)
	time.Sleep(5 * time.Second) // 5s 后终止
}
