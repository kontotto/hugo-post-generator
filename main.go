package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Metadata struct {
	Categories  []string `json:"categories"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Url         string   `json:"url"`
}

func main() {
	bytes, err := ioutil.ReadFile("movies/20181120/metadata.json")
	if err != nil {
		log.Fatal(err)
	}

	var metadata Metadata
	if err := json.Unmarshal(bytes, &metadata); err != nil {
		log.Fatal(err)
	}

	fmt.Println("%v", metadata)
}
