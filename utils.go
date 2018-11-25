package main

import (
	"io/ioutil"
	"path/filepath"
)

type DateFile struct {
	Date     string
	Filename string
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
