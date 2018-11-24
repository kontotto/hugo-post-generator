package main

import (
	"testing"
)

func TestOpenTemplateSuccess(t *testing.T) {
	tmpl := OpenTemplate("tests/movies/20100101/sm12345678.md.tmpl")

	want := "sm12345678.md.tmpl"
	got := tmpl.Name()

	if want != got {
		t.Fatalf("want = %s. got = %s", want, got)
	}
}

func TestBuildPostTimeSuccess(t *testing.T) {
	want := "2010-01-01T19:00:00Z"
	got, err := BuildPostTime("2010-01-01")

	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatalf("want = %s, got = %s", want, got)
	}
}

func TestBuildPostTimeFail(t *testing.T) {
	_, err := BuildPostTime("2010/01/01")

	if err == nil {
		t.Fatal()
	}
}
