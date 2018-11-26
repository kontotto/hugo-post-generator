package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

var generateCommand = cli.Command{
	Name:      "generate",
	Usage:     "generate a hugo post",
	ArgsUsage: `<templates-path> <hugosite-path>`,
	Action: func(context *cli.Context) error {
		tmpls, err := OpenTemplates(context.Args().Get(0))
		if err != nil {
			return err
		}

		for _, tmpl := range tmpls {
			meta, err := NewMeta(tmpl.Date, tmpl.Filename, context.Args().Get(1))
			if err != nil {
				return err
			}

			factory, err := NewFactory(meta)
			if err != nil {
				return err
			}

			provider, err := factory.CreateProvider()
			if err != nil {
				return err
			}

			data, err := provider.Data()
			if err != nil {
				return err
			}

			path := filepath.Join(context.Args().Get(1), "/content/posts/"+meta.Id+".md")
			file, err := os.Create(path)
			if err != nil {
				return err
			}

			if err := tmpl.Template.Execute(file, data); err != nil {
				return err
			}

			fmt.Println(path, " created")
		}
		return nil
	},
}
