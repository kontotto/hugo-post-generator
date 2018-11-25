package main

type MovieProvider interface {
	Provider
	Category() string
	Date() string
	Thumbnail() string
	Title() string
	Embed() string
}
