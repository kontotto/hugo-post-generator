package main

import (
	"text/template"
)

type Provider interface {
	Template() *template.Template
}
