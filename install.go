package main

import (
	"context"
	"github.com/ngyewch/mise-helper/helper"
	"github.com/urfave/cli/v3"
)

func install(ctx context.Context, cmd *cli.Command) error {
	recursive := cmd.Bool(recursiveFlag.Name)

	return helper.Install(recursive)
}
