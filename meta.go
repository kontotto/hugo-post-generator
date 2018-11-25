package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

type Meta struct {
	Date     string
	Id       string
	HugoPath string
}

func NewMeta(date string, filename string, hugopath string) (*Meta, error) {
	id, err := getId(filename)
	if err != nil {
		return nil, err
	}

	abspath, err := filepath.Abs(hugopath)
	if err != nil {
		return nil, err
	}

	return &Meta{
		Date:     date,
		Id:       id,
		HugoPath: abspath,
	}, nil
}

func getId(filename string) (string, error) {
	matched, err := regexp.MatchString(`\.md\.tmpl`, filename)
	if err != nil {
		return "", err
	}

	if !matched {
		return "", fmt.Errorf("%s is unmatched", filename)
	}

	return strings.Trim(filename, `\.md\.tmpl`), nil
}
