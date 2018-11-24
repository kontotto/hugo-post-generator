package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
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

// sm12345678.md.tmpl
func BuildCategory(filename string) (string, error) {
	r := regexp.MustCompile(`^sm[0-9]{1,8}`)
	if !r.MatchString(filename) {
		return "", fmt.Errorf("%s is unmatched", filename)
	}

	return "niconico", nil
}

// 2000-01-01 ~ 2029-12-31
// TODO: more correctly
func BuildPostTime(date string) (string, error) {
	r := regexp.MustCompile(`^20[0-2][0-9]-[0|1][0-9]-[0-3][0-9]$`)
	if !r.MatchString(date) {
		return "", fmt.Errorf("%s is unmatched", date)
	}

	return date + "T19:00:00Z", nil
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
