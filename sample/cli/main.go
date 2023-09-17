package main

import (
	"log"
	"os"

	"github.com/onnga-wasabi/go-ddd/sample/cli/command"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			command.UserSandboxCommand(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
