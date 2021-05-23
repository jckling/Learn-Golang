package main

import (
	"./rpc_objects"
	"fmt"
	"log"
	"net/rpc"
)

const serverAddress = "localhost"

// go run rpc_client.go
func main() {
	client, err := rpc.DialHTTP("tcp", serverAddress+":8080")
	if err != nil {
		log.Fatal("Error dialing:", err)
	}

	// Synchronous call
	args := &rpc_objects.Args{7, 8} // 初始化对象
	var reply int                   // 存储调用结果

	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Args error:", err)
	}

	fmt.Printf("Args: %d * %d = %d", args.N, args.M, reply)
}

/* Output:
Args: 7 * 8 = 56
*/
