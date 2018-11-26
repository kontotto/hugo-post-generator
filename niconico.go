package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const imageUrl = "http://tn-skr3.smilevideo.jp/smile?i="

type niconicoProvider struct {
	Meta *Meta
}

func (p *niconicoProvider) Data() (interface{}, error) {
	thumbnail, err := p.Thumbnail()
	if err != nil {
		return nil, err
	}

	title, err := p.Title()
	if err != nil {
		return nil, err
	}

	return &MovieData{
		Category:  p.Category(),
		Date:      p.Date(),
		Thumbnail: thumbnail,
		Title:     title,
		Embed:     p.Embed(),
	}, nil
}

func (p *niconicoProvider) Category() string {
	return "niconico"
}

func (p *niconicoProvider) Date() string {
	return p.Meta.Date
}

func (p *niconicoProvider) Thumbnail() (string, error) {
	numberId, err := idExceptPrefix(p.Meta.Id)
	if err != nil {
		return "", err
	}
	jpgUrl := imageUrl + numberId + ".L"
	storePath := "./tests/static/images/" + p.Meta.Id + ".jpg"

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

func (p *niconicoProvider) Title() (string, error) {
	doc, err := goquery.NewDocument("https://www.nicovideo.jp/watch/" + p.Meta.Id)
	if err != nil {
		return "", err
	}

	return doc.Find("title").First().Text(), nil
}

func (p *niconicoProvider) Embed() string {
	return ""
}

func idExceptPrefix(id string) (string, error) {
	matched, err := regexp.MatchString(`^sm`, id)
	if err != nil {
		return "", err
	}

	if !matched {
		return "", fmt.Errorf("%s is unmatched", id)
	}

	return strings.Trim(id, `^sm`), nil
}
