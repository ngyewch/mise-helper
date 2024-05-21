package cmd

import (
	"github.com/urfave/cli/v2"
)

var (
	app = &cli.App{
		Name:   "mise-helper",
		Usage:  "mise helper",
		Action: nil,
		Commands: []*cli.Command{
			installCmd,
			latestCmd,
		},
	}
)

func Run(args []string) error {
	return app.Run(args)
}
