package main

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetTimeSuccess(t *testing.T) {
	want := "2010-01-01T19:00:00Z"
	got, err := getTime("2010-01-01")

	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestGetIdSuccess(t *testing.T) {
	want := "sm22222222"
	got, err := getId("sm22222222.md.tmpl")

	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestNewMetaSuccess(t *testing.T) {
	abspath, err := filepath.Abs("./tests")
	if err != nil {
		t.Fatal(err)
	}

	want := &Meta{
		Time:     "2010-01-01T19:00:00Z",
		Id:       "sm22222222",
		HugoPath: abspath,
	}
	got, err := NewMeta("2010-01-01", "sm22222222.md.tmpl", "./tests")

	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}
