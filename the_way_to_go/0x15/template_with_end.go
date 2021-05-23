package main

import (
	"os"
	"text/template"
)

// {{.}} 被设置为当前管道的值
// with 语句将点号设为管道的值
// 如果管道是空的，那么不管 with-end 块之间有什么，都会被忽略
// with-end
func main() {
	t := template.New("test")
	t, _ = t.Parse("{{with `hello`}}{{.}}{{end}}!\n")
	t.Execute(os.Stdout, nil)

	t, _ = t.Parse("{{with `hello`}}{{.}} {{with `Mary`}}{{.}}{{end}}{{end}}!\n")
	t.Execute(os.Stdout, nil)
}
