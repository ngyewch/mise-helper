package cmd

import (
	versionInfoCobra "github.com/ngyewch/go-versioninfo/cobra"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:  AppName,
		RunE: help,
	}
)

func help(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	versionInfoCobra.AddVersionCmd(rootCmd, nil)
}
