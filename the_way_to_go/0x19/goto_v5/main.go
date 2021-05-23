package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/rpc"
)

var (
	listenAddr = flag.String("http", ":8080", "http listen address")
	dataFile   = flag.String("file", "store.gob", "data store file name")
	hostname   = flag.String("host", "localhost:8080", "http host name")
	masterAddr = flag.String("master", "", "RPC master address")
	rpcEnabled = flag.Bool("rpc", false, "enable RPC server") // 是否开启 rpc
)

var store Store

func main() {
	flag.Parse() // 解析命令行参数
	if *masterAddr != "" {
		store = NewProxyStore(*masterAddr) // slave
	} else {
		store = NewURLStore(*dataFile) // master
	}
	if *rpcEnabled { // the master is the rpc server:
		rpc.RegisterName("Store", store) // 注册 rpc
		rpc.HandleHTTP()
	}
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	http.ListenAndServe(*listenAddr, nil)

}

func Redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	if key == "" {
		http.NotFound(w, r)
		return
	}
	var url string
	if err := store.Get(&key, &url); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func Add(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, AddForm)
		return
	}
	var key string
	if err := store.Put(&url, &key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "http://%s/%s", *hostname, key)
}

const AddForm = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`

/*
goto master: go run key.go store.go main.go -http=:8081 -rpc=true
goto slave: go run key.go store.go main.go -master=127.0.0.1:8081
*/
