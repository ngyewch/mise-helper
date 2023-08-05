package cmd

import (
	"github.com/ngyewch/rtx-helper/helper"
	"github.com/spf13/cobra"
)

var (
	installCmd = &cobra.Command{
		Use:  "install",
		RunE: install,
	}
)

func install(cmd *cobra.Command, args []string) error {
	return helper.Install()
}

func init() {
	rootCmd.AddCommand(installCmd)
}
