package main

import (
	"fmt"
	"log"
	"text/template"
)

func main() {
	tOk := template.New("ok")
	tErr := template.New("error_template")
	defer func() {
		if err := recover(); err != nil { // 处理错误
			log.Printf("run time panic: %v", err)
		}
	}()

	//a valid template, so no panic with Must:
	template.Must(tOk.Parse("/* and a comment */ some static text: {{ .Name }}")) // 正确模板
	fmt.Println("The first one parsed OK.")
	fmt.Println("The next one ought to fail.")
	template.Must(tErr.Parse(" some static text {{ .Name }")) // 错误模板
}
