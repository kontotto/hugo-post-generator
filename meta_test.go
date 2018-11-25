package main

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetIdSuccess(t *testing.T) {
	want := "sm12345678"
	got, err := getId("sm12345678.md.tmpl")

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
		Date:     "2010-01-01",
		Id:       "sm12345678",
		HugoPath: abspath,
	}
	got, err := NewMeta("2010-01-01", "sm12345678.md.tmpl", "./tests")

	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}
