package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func create() (MovieProvider, error) {
	meta, err := NewMeta("2010-01-01", "sm22222222.md.tmpl", "./tests")
	if err != nil {
		return nil, err
	}

	return &niconicoProvider{
		Meta: meta,
	}, nil
}

func TestNiconicoData(t *testing.T) {
	provider, err := create()
	if err != nil {
		t.Fatal(err)
	}

	abspath, err := filepath.Abs("./tests/static/images/sm22222222.jpg")
	if err != nil {
		t.Fatal(err)
	}

	want := &MovieData{
		Category:  "niconico",
		Date:      "2010-01-01",
		Thumbnail: abspath,
		Title:     "ある日の鎌倉の風景 - ニコニコ動画",
		Embed:     "",
	}
	got, err := provider.Data()
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Fatalf("want %+v, got %+v", want, got)
	}
}

func TestNiconicoCategory(t *testing.T) {
	provider, err := create()
	if err != nil {
		t.Fatal(err)
	}

	want := "niconico"
	got := provider.Category()

	if !cmp.Equal(want, got) {
		t.Fatalf("want %+v, got %+v", want, got)
	}
}

func TestNiconicoDate(t *testing.T) {
	provider, err := create()
	if err != nil {
		t.Fatal(err)
	}

	want := "2010-01-01"
	got := provider.Date()

	if !cmp.Equal(want, got) {
		t.Fatalf("want %+v, got %+v", want, got)
	}
}

func TestNiconicoThumbnail(t *testing.T) {
	provider, err := create()
	if err != nil {
		t.Fatal(err)
	}

	abspath, err := filepath.Abs("./tests/static/images/sm22222222.jpg")
	if err != nil {
		t.Fatal(err)
	}

	want := abspath
	got, err := provider.Thumbnail()
	if err != nil {
		t.Fatal(err)
	}

	if err := os.Remove("./tests/static/images/sm22222222.jpg"); err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Fatalf("want %+v, got %+v", want, got)
	}
}

func TestNiconicoTitle(t *testing.T) {
	provider, err := create()
	if err != nil {
		t.Fatal(err)
	}

	want := "ある日の鎌倉の風景 - ニコニコ動画"
	got, err := provider.Title()
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Fatalf("want %+v, got %+v", want, got)
	}
}

func TestIdExceptPrefixSuccess(t *testing.T) {
	want := "22222222"
	got, err := idExceptPrefix("sm22222222")
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Fatalf("want %+v, got %+v", want, got)
	}
}

func TestIdExceptPrefixFail(t *testing.T) {
	_, err := idExceptPrefix("ssm22222222")
	if err == nil {
		t.Fatal()
	}
}
