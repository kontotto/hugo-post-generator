package main

type MovieProvider interface {
	Provider
	Category() string
	Date() string
	Thumbnail() (string, error)
	Title() (string, error)
	Embed() (string, error)
}

type MovieData struct {
	Category  string
	Date      string
	Thumbnail string
	Title     string
	Embed     string
}
