package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewFactoryNiconico(t *testing.T) {
	meta := &Meta{
		Date: "2010-01-01",
		Id:   "sm12345678",
	}

	want := &niconicoFactory{
		Meta: meta,
	}

	got, err := NewFactory(meta)

	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Fatalf("want %+v, got %+v", want, got)
	}
}
