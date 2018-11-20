package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type Metadata struct {
	Categories  []string `json:"categories"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Url         string   `json:"url"`
}

type Template struct {
	Categories  []string
	Date        string
	Description string
	Movie       string
	Tags        []string
	Thumbnail   string
}

func BuildTemplate(m *Metadata) (*Template, error) {
	t := &Template{
		Categories:  m.Categories,
		Date:        "",
		Description: m.Description,
		Movie:       "",
		Tags:        m.Tags,
		Thumbnail:   "",
	}
	return t, nil
}

func OpenTemplate(path string) (*template.Template, error) {
	t := template.Must(template.ParseFiles(path))
	return t, nil
}

func OpenMetadata(path string) (*Metadata, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	metadata := &Metadata{}
	if err := json.Unmarshal(bytes, &metadata); err != nil {
		return nil, err
	}
	return metadata, nil
}

func main() {
	tmpl, err := OpenTemplate("tests/testdata/templates/movie.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	m, err := OpenMetadata("tests/testdata/movies/20181120/metadata.json")
	if err != nil {
		log.Fatal(err)
	}

	t, err := BuildTemplate(m)
	if err != nil {
		log.Fatal(err)
	}

	if err := tmpl.Execute(os.Stdout, t); err != nil {
		log.Fatal(err)
	}

}
