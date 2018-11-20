package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Metadata struct {
	Categories  []string `json:"categories"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Url         string   `json:"url"`
}

func OpenMetadata(path string) (*Metadata, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	metadata := &Metadata{}
	if err := json.Unmarshal(bytes, &metadata); err != nil {
		return nil, err
	}
	return metadata, nil
}

func main() {
	fmt.Println("Hello World")
}
