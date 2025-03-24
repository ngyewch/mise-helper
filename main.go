package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var (
	version string

	recursiveFlag = &cli.BoolFlag{
		Name:  "recursive",
		Usage: "run recursively",
		Value: true,
	}
	hideLatestFlag = &cli.BoolFlag{
		Name:  "hide-latest",
		Usage: "do not print tools already at the latest version",
		Value: true,
	}
	includePrereleasesFlag = &cli.BoolFlag{
		Name:  "include-prereleases",
		Usage: "include prereleases",
		Value: false,
	}
	includeGlobalFlag = &cli.BoolFlag{
		Name:  "include-global",
		Usage: "include global",
		Value: false,
	}

	installCmd = &cli.Command{
		Name:   "install",
		Usage:  "Install recursively",
		Action: install,
		Flags: []cli.Flag{
			recursiveFlag,
		},
	}
	latestCmd = &cli.Command{
		Name:   "latest",
		Usage:  "Show latest versions recursively",
		Action: latest,
		Flags: []cli.Flag{
			recursiveFlag,
			hideLatestFlag,
			includePrereleasesFlag,
			includeGlobalFlag,
		},
	}

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
