package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

type Meta struct {
	Time     string
	Id       string
	HugoPath string
}

func NewMeta(date string, filename string, hugopath string) (*Meta, error) {
	time, err := getTime(date)
	if err != nil {
		return nil, err
	}

	id, err := getId(filename)
	if err != nil {
		return nil, err
	}

	abspath, err := filepath.Abs(hugopath)
	if err != nil {
		return nil, err
	}

	return &Meta{
		Time:     time,
		Id:       id,
		HugoPath: abspath,
	}, nil
}

func getTime(date string) (string, error) {
	r := regexp.MustCompile(`^20[0-2][0-9]-[0|1][0-9]-[0-3][0-9]$`)
	if !r.MatchString(date) {
		return "", fmt.Errorf("%s is unmatched", date)
	}

	return date + "T19:00:00Z", nil
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
