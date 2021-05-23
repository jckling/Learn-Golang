package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (this *Page) save() (err error) {
	return ioutil.WriteFile(this.Title, this.Body, 0666)
}

func (this *Page) load(title string) (err error) {
	this.Title = title
	this.Body, err = ioutil.ReadFile(this.Title)
	return err
}

func main() {
	page := Page{
		"D:/Study/Github/Learn-Golang/the_way_to_go/0x12/Page.md",
		[]byte("# Page\n## Section1\nThis is section1."),
	}
	page.save()

	// load from Page.md
	var new_page Page
	new_page.load("D:/Study/Github/Learn-Golang/the_way_to_go/0x12/Page.md")
	fmt.Println(string(new_page.Body))
}
