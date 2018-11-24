package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Meta struct {
	Date string
	Id   string
}

func NewMeta(date string, filename string) (*Meta, error) {
	id, err := getId(filename)
	if err != nil {
		return nil, err
	}

	return &Meta{
		Date: date,
		Id:   id,
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
