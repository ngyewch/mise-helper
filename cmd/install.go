package cmd

import (
	"github.com/ngyewch/mise-helper/helper"
	"github.com/urfave/cli/v2"
)

var (
	installCmd = &cli.Command{
		Name:   "install",
		Usage:  "Install recursively",
		Action: install,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "recursive",
				Usage: "run recursively",
				Value: true,
			},
		},
	}
)

func install(cCtx *cli.Context) error {
	recursive := cCtx.Bool("recursive")

	return helper.Install(recursive)
}
