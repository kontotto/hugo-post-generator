package main

import (
	"log"
	"os"
	"text/template"
)

type Template struct {
	Category  string
	Date      string
	Embed     string
	Thumbnail string
	Title     string
}

func BuildTemplate() (*Template, error) {
	t := &Template{
		Category:  "",
		Date:      "",
		Embed:     "",
		Thumbnail: "",
		Title:     "",
	}
	return t, nil
}

func OpenTemplate(path string) *template.Template {
	return template.Must(template.ParseFiles(path))
}

func main() {
	tmpl := OpenTemplate("tests/movies/20100101/sm12345678.md.tmpl")

	t, err := BuildTemplate()
	if err != nil {
		log.Fatal(err)
	}

	if err := tmpl.Execute(os.Stdout, t); err != nil {
		log.Fatal(err)
	}

}
