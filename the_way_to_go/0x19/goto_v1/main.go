package main

import (
	"fmt"
	"net/http"
)

// 表单
const AddForm = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`

var store = NewURLStore()

// go run key.go store.go main.go
func main() {
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	http.ListenAndServe(":8080", nil)
}

// 短网址重定向
func Redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:] // 键
	url := store.Get(key) // 查找对应的长网址
	if url == "" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

// 提交新网址
func Add(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url") // 读取长网址
	if url == "" {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, AddForm) // 输出表单
		return
	}
	key := store.Put(url)                           // 生成短网址
	fmt.Fprintf(w, "http://localhost:8080/%s", key) // 发送短网址
}

/*
1. 添加网址：http://localhost:8080/add
2. 提交表单：http://golang.org/pkg/bufio/#Writer
3. 访问短网址：http://localhost:8080/0
*/
