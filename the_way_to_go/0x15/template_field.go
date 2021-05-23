package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name                string
	nonExportedAgeField string
}

// {{.FieldName}}
func main() {
	t := template.New("hello")
	t, _ = t.Parse("hello {{.Name}}!") // 解析模板，定义字符串
	p := Person{Name: "Mary", nonExportedAgeField: "31"}
	if err := t.Execute(os.Stdout, p); err != nil { // 字段替换
		fmt.Println("There was an error:", err.Error())
	}
}
