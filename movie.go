package main

type MoviePorivder interface {
	Provider
	Category() string
	Date() string
	Thumbnail() string
	Title() string
	Embed() string
}
