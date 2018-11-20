package main

import (
	"testing"
)

func TestOpenMetadataSuccess(t *testing.T) {
	_, err := OpenMetadata("tests/testdata/movies/20181120/metadata.json")
	if err != nil {
		t.Fatal("failed test")
	}
}

func TestOpenMetadataFailed(t *testing.T) {
	_, err := OpenMetadata("unknown.json")
	if err == nil {
		t.Fatal("failed test")
	}
}

func TestOpenTemplateSuccess(t *testing.T) {
	_, err := OpenTemplate("tests/testdata/templates/movie.tmpl")
	if err != nil {
		t.Fatal("failed test")
	}
}
