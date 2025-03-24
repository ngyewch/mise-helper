package main

import (
	"github.com/ngyewch/mise-helper/helper"
	"github.com/urfave/cli/v2"
)

func install(cCtx *cli.Context) error {
	recursive := recursiveFlag.Get(cCtx)

	return helper.Install(recursive)
}
