package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var (
	version string

	app = &cli.App{
		Name:    "mise-helper",
		Usage:   "mise helper",
		Version: version,
		Commands: []*cli.Command{
			installCmd,
			latestCmd,
		},
	}
)

func main() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
