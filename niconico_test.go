package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func create() MovieProvider {
	meta := &Meta{
		Date: "2010-01-01",
		Id:   "sm12345678",
	}
	return &niconicoProvider{
		Meta: meta,
	}
}

func TestNiconicoCategory(t *testing.T) {
	provider := create()

	want := "niconico"
	got := provider.Category()

	if !cmp.Equal(want, got) {
		t.Fatalf("want %+v, got %+v", want, got)
	}
}
