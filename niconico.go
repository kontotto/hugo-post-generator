package main

import (
	"text/template"
)

type niconicoProvider struct {
	Meta *Meta
}

func (p *niconicoProvider) Template() *template.Template {
	return nil
}

func (p *niconicoProvider) Category() string {
	return "niconico"
}

func (p *niconicoProvider) Date() string {
	return p.Meta.Date
}

func (p *niconicoProvider) Thumbnail() string {
	return ""
}

func (p *niconicoProvider) Title() string {
	return ""
}

func (p *niconicoProvider) Embed() string {
	return ""
}
