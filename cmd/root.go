// Package cmd contains all CLI commands used by the application.
package cmd

import (
	"fmt"

	"github.com/caarlos0/log"
	"github.com/spf13/cobra"
)

// application build information set by the linker.
var (
	version  string
	revision string
)

var rootCmd *cobra.Command

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := GetRootCommand().Execute()
	if err != nil {
		log.WithError(err).Fatal("fatal error")
	}
}

// GetRootCommand returns the root command for the CLI.
func GetRootCommand() *cobra.Command {
	if rootCmd == nil {
		rootCmd = &cobra.Command{
			Use:     "newed",
			Version: fmt.Sprintf("%s (%s)", version, revision),
			Short:   "newed creates projects from templates",
			Args:    cobra.NoArgs,
			Run: func(cmd *cobra.Command, _ []string) {
				_ = cmd.Help()
			},
		}

		rootCmd.PersistentFlags().StringP("config", "c", "", "configuration file")

		rootCmd.AddCommand(GetApplyCmd())
		rootCmd.AddCommand(GetListCmd())
		rootCmd.AddCommand(GetVersionCmd())
	}

	return rootCmd
}
