package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "hugo-post-generator"
	app.Usage = "generate hugo post from template file"

	app.Commands = []cli.Command{
		generateCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
