package main

import (
	"context"
	"github.com/ngyewch/mise-helper/helper"
	"github.com/urfave/cli/v3"
)

func latest(ctx context.Context, cmd *cli.Command) error {
	hideLatest := cmd.Bool(hideLatestFlag.Name)
	includePrereleases := cmd.Bool(includePrereleasesFlag.Name)
	recursive := cmd.Bool(recursiveFlag.Name)
	includeGlobal := cmd.Bool(includeGlobalFlag.Name)

	return helper.Latest(hideLatest, includePrereleases, recursive, includeGlobal)
}
