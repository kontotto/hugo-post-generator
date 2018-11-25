package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

const imageUrl = "http://tn-skr3.smilevideo.jp/smile?i="

type niconicoProvider struct {
	Meta *Meta
}

func (p *niconicoProvider) Template() *template.Template {
	return nil
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

func (p *niconicoProvider) Title() string {
	return ""
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
