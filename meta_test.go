package main

import (
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
	want := &Meta{
		Date: "2010-01-01",
		Id:   "sm12345678",
	}
	got, err := NewMeta("2010-01-01", "sm12345678.md.tmpl")

	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}
