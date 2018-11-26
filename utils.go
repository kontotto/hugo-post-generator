package main

import (
	"io/ioutil"
	"path/filepath"
	"text/template"
)

type DateFile struct {
	Date     string
	Filename string
}

type MetaTemplate struct {
	DateFile
	Template *template.Template
}

func GetTemplatePaths(tmpDir string) ([]DateFile, error) {
	dirs, err := ioutil.ReadDir(tmpDir)
	if err != nil {
		return nil, err
	}

	var datefiles []DateFile
	for _, dir := range dirs {
		files, err := ioutil.ReadDir(filepath.Join(tmpDir, dir.Name()))
		if err != nil {
			return nil, err
		}

		for _, file := range files {
			tmp := DateFile{
				Date:     dir.Name(),
				Filename: file.Name(),
			}
			datefiles = append(datefiles, tmp)
		}
	}

	return datefiles, nil
}

func OpenTemplates(tmpDir string) ([]MetaTemplate, error) {
	datefiles, err := GetTemplatePaths(tmpDir)
	if err != nil {
		return nil, err
	}

	var templates []MetaTemplate
	for _, datefile := range datefiles {
		path := filepath.Join(tmpDir, datefile.Date, datefile.Filename)

		tmpl := MetaTemplate{
			DateFile: datefile,
			Template: template.Must(template.ParseFiles(path)),
		}
		templates = append(templates, tmpl)
	}
	return templates, nil
}
