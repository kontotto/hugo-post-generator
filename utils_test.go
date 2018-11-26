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

func TestOpenTemplates(t *testing.T) {
	want := []MetaTemplate{
		MetaTemplate{
			DateFile: DateFile{
				Date:     "2010-01-01",
				Filename: "sm12345678.md.tmpl",
			},
			// TODO: correct test
			Template: nil,
		},
	}
	got, err := OpenTemplates("./tests/movies/")
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want[0].DateFile, got[0].DateFile) {
		t.Fatalf("want %v, got %v", want, got)
	}
}
