package main

import (
	"github.com/ngyewch/mise-helper/helper"
	"github.com/urfave/cli/v2"
)

func latest(cCtx *cli.Context) error {
	hideLatest := hideLatestFlag.Get(cCtx)
	includePrereleases := includePrereleasesFlag.Get(cCtx)
	recursive := recursiveFlag.Get(cCtx)
	includeGlobal := includeGlobalFlag.Get(cCtx)

	return helper.Latest(hideLatest, includePrereleases, recursive, includeGlobal)
}
