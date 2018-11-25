package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetTemplatePaths(t *testing.T) {
	want := []DateFile{
		DateFile{
			Date:     "2010-01-01",
			Filename: "sm12345678.md.tmpl",
		},
	}
	got, err := GetTemplatePaths("./tests/movies/")
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}
