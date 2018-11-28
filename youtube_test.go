package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func createYoutubeProvider() (MovieProvider, error) {
	meta, err := NewMeta("2015-01-01", "Ct6BUPvE2sM.md.tmpl", "./tests")
	if err != nil {
		return nil, err
	}

	return &youtubeProvider{
		Meta: meta,
	}, nil
}

func TestYoutubeData(t *testing.T) {
	provider, err := createYoutubeProvider()
	if err != nil {
		t.Fatal(err)
	}

	abspath, err := filepath.Abs("./tests/static/images/Ct6BUPvE2sM.jpg")
	if err != nil {
		t.Fatal(err)
	}

	want := &MovieData{
		Category:  "youtube",
		Date:      "2015-01-01T19:00:00Z",
		Thumbnail: abspath,
		Title:     "PIKOTARO - PPAP (Pen Pineapple Apple Pen) (Long Version) [Official Video] - YouTube",
		Embed:     `<iframe width="560" height="315" src="https://www.youtube.com/embed/Ct6BUPvE2sM" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>`,
	}
	got, err := provider.Data()
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Fatalf("want %+v, got %+v", want, got)
	}
}

func TestYoutubeCategory(t *testing.T) {
	provider, err := createYoutubeProvider()
	if err != nil {
		t.Fatal(err)
	}

	want := "youtube"
	got := provider.Category()

	if !cmp.Equal(want, got) {
		t.Fatalf("want %+v, got %+v", want, got)
	}
}

func TestYoutubeDate(t *testing.T) {
	provider, err := createYoutubeProvider()
	if err != nil {
		t.Fatal(err)
	}

	want := "2015-01-01T19:00:00Z"
	got := provider.Date()

	if !cmp.Equal(want, got) {
		t.Fatalf("want %+v, got %+v", want, got)
	}
}

func TestYoutubeThumbnail(t *testing.T) {
	provider, err := createYoutubeProvider()
	if err != nil {
		t.Fatal(err)
	}

	abspath, err := filepath.Abs("./tests/static/images/Ct6BUPvE2sM.jpg")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove("./tests/static/images/Ct6BUPvE2sM.jpg")

	want := abspath
	got, err := provider.Thumbnail()
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Fatalf("want %+v, got %+v", want, got)
	}
}

func TestYoutubeTitle(t *testing.T) {
	provider, err := createYoutubeProvider()
	if err != nil {
		t.Fatal(err)
	}

	want := "PIKOTARO - PPAP (Pen Pineapple Apple Pen) (Long Version) [Official Video] - YouTube"
	got, err := provider.Title()
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Fatalf("want %+v, got %+v", want, got)
	}
}

func TestYoutubeEmbed(t *testing.T) {
	provider, err := createYoutubeProvider()
	if err != nil {
		t.Fatal(err)
	}

	want := `<iframe width="560" height="315" src="https://www.youtube.com/embed/Ct6BUPvE2sM" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>`
	got, err := provider.Embed()
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Fatalf("want %+v, got %+v", want, got)
	}
}
