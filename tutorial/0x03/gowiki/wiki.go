package main

import (
	"fmt"
	"html/template" // process HTML templates
	"io/ioutil"     // save/read text file
	"log"
	"net/http" // build web applications
	"regexp"   // validate user input
)

// The Page struct describes how page data will be stored in memory.
type Page struct {
	Title string // 标题
	Body  []byte // 页面内容，字节切片，提供给 io 包的类型
}

// save the Page's Body to a text file.
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// loadPage returns a pointer to a Page literal constructed with the proper title and body values.
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// viewHandler will allow users to view a wiki page.
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

// editHandler loads the page (or, if it doesn't exist, create an empty Page struct), and displays an HTML form.
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

// saveHandler will handle the submission of forms located on the edit pages.
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

// create a global variable named templates, and initialize it with ParseFiles.
// The ParseFiles function takes any number of string arguments that identify our template files,
// and parses those files into templates that are named after the base file name.
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p) // 渲染模板
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// create a global variable validPath to storevalidation expression
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// a wrapper function that takes a function of the above type, and returns a function of type http.HandlerFunc
// The returned function is called a closure because it encloses values defined outside of it.
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	// save to and load from a file.
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))

	// wrap the handler functions with makeHandler, before they are registered with the http package
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
view existed page named "TestPage"
http://localhost:8080/view/TestPage

create new page named "test"
http://localhost:8080/view/test
*/
