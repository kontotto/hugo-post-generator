package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewFactoryNiconico(t *testing.T) {
	meta, err := NewMeta("2010-01-01", "sm22222222.md.tmpl", "./tests")
	if err != nil {
		t.Fatal(err)
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
