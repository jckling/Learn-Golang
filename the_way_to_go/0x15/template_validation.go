package main

import (
	"fmt"
	"text/template"
)

// 验证模板格式
func main() {
	tOk := template.New("ok")
	//a valid template, so no panic with Must:
	template.Must(tOk.Parse("/* and a comment */ some static text: {{ .Name }}")) // 正确模板

	fmt.Println("The first one parsed OK.")
	fmt.Println("The next one ought to fail.")

	tErr := template.New("error_template")
	template.Must(tErr.Parse(" some static text {{ .Name }")) // 错误模板
}
