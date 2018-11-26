package main

import (
	"log"
	"os"
)

func main() {
	tmpls, err := OpenTemplates("./tests/movies")
	if err != nil {
		log.Fatal(err)
	}

	for _, tmpl := range tmpls {
		meta, err := NewMeta(tmpl.Date, tmpl.Filename, "./tests")
		if err != nil {
			log.Fatal(err)
		}

		factory, err := NewFactory(meta)
		if err != nil {
			log.Fatal(err)
		}

		provider, err := factory.CreateProvider()
		if err != nil {
			log.Fatal(err)
		}

		data, err := provider.Data()
		if err != nil {
			log.Fatal(err)
		}

		if err := tmpl.Template.Execute(os.Stdout, data); err != nil {
			log.Fatal(err)
		}
	}
}
