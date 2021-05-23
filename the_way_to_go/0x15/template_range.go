package main

import (
	"os"
	"text/template"
)

// {{range pipeline}} T1 {{else}} T0 {{end}}
// 管道的值必须是数组、切片或 map
// 如果管道的值长度为零，点号的值不受影响，且执行 T0；否则，点号被设置为数组、切片或 map 内元素的值，并执行 T1。
// range-end
func main() {
	t := template.New("template test")
	t.Parse("{{range .}}{{.}}{{end}}")
	s := []int{1, 2, 3, 4}
	t.Execute(os.Stdout, s)
}
