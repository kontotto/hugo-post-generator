package main

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type youtubeProvider struct {
	*Meta
}

func (p *youtubeProvider) Data() (interface{}, error) {
	thumbnail, err := p.Thumbnail()
	if err != nil {
		return nil, err
	}

	title, err := p.Title()
	if err != nil {
		return nil, err
	}

	embed, err := p.Embed()
	if err != nil {
		return nil, err
	}

	return &MovieData{
		Category:  p.Category(),
		Date:      p.Date(),
		Thumbnail: thumbnail,
		Title:     title,
		Embed:     embed,
	}, nil
}

func (p *youtubeProvider) Category() string {
	return "youtube"
}

func (p *youtubeProvider) Date() string {
	return p.Time
}

func (p *youtubeProvider) Thumbnail() (string, error) {
	jpgUrl := "https://i.ytimg.com/vi/" + p.Id + "/default.jpg"
	storePath := p.HugoPath + "/static/images/" + p.Id + ".jpg"

	jpg, err := os.Create(storePath)
	if err != nil {
		return "", err
	}
	defer jpg.Close()

	res, err := http.Get(jpgUrl)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	_, err = io.Copy(jpg, res.Body)
	if err != nil {
		return "", err
	}

	abspath, err := filepath.Abs(storePath)
	if err != nil {
		return "", nil
	}

	return abspath, nil
}

func (p *youtubeProvider) Title() (string, error) {
	doc, err := goquery.NewDocument("https://www.youtube.com/watch?v=" + p.Id)
	if err != nil {
		return "", err
	}

	return doc.Find("title").First().Text(), nil
}

func (p *youtubeProvider) Embed() (string, error) {
	return `<iframe width="560" height="315" src="https://www.youtube.com/embed/` + p.Id + `" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>`, nil
}
