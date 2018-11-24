package main

import (
	"testing"
)

func TestOpenTemplateSuccess(t *testing.T) {
	tmpl := OpenTemplate("tests/movies/20100101/sm12345678.md.tmpl")

	expected := "sm12345678.md.tmpl"
	actual := tmpl.Name()

	if expected != actual {
		t.Fail()
	}
}
