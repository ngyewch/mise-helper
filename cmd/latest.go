package cmd

import (
	"github.com/ngyewch/mise-helper/helper"
	"github.com/urfave/cli/v2"
)

var (
	latestCmd = &cli.Command{
		Name:   "latest",
		Usage:  "Show latest versions recursively",
		Action: latest,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "hide-latest",
				Usage: "do not print tools already at the latest version",
				Value: true,
			},
			&cli.BoolFlag{
				Name:  "include-prereleases",
				Usage: "include prereleases",
				Value: false,
			},
			&cli.BoolFlag{
				Name:  "recursive",
				Usage: "run recursively",
				Value: true,
			},
		},
	}
)

func latest(cCtx *cli.Context) error {
	hideLatest := cCtx.Bool("hide-latest")
	includePrereleases := cCtx.Bool("include-prereleases")
	recursive := cCtx.Bool("recursive")

	return helper.Latest(hideLatest, includePrereleases, recursive)
}
