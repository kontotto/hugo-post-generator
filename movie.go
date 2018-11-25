package main

type MovieProvider interface {
	Provider
	Category() string
	Date() string
	Thumbnail() (string, error)
	Title() string
	Embed() string
}
