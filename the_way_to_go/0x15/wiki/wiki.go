package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"text/template"
)

const lenPath = len("/view/")

var titleValidator = regexp.MustCompile("^[a-zA-Z0-9]+$") // 正则检查
var templates = make(map[string]*template.Template)       // 映射索引（模板缓存）
var err error

type Page struct {
	Title string // 标题
	Body  []byte // 文本内容
}

func init() {
	for _, tmpl := range []string{"edit", "view"} {
		templates[tmpl] = template.Must(template.ParseFiles(tmpl + ".html")) // 模板文件转换成 *template.Template 对象
	}
}

// c
func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler)) // http://localhost:8080/view
	http.HandleFunc("/edit/", makeHandler(editHandler)) // http://localhost:8080/edit
	http.HandleFunc("/save/", makeHandler(saveHandler)) // http://localhost:8080/save
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

// 返回闭包函数
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Path[lenPath:]
		if !titleValidator.MatchString(title) { // 验证标题
			http.NotFound(w, r)
			return
		}
		fn(w, r, title) // 构造返回值
	}
}

// 读取文件
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := load(title)
	if err != nil { // page not found
		http.Redirect(w, r, "/edit/"+title, http.StatusFound) // 没有则创建
		return
	}
	renderTemplate(w, "view", p)
}

// 编辑文件
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := load(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

// 保存文件
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// 渲染模板
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates[tmpl].Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// 保存文件
func (p *Page) save() error {
	filename := p.Title + ".txt"
	// file created with read-write permissions for the current user only
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func load(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
