// Package cmd contains all CLI commands used by the application.
package cmd

import (
	"fmt"
	"os"

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
		os.Exit(1)
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
		}

		rootCmd.PersistentFlags().StringP("config", "c", "", "configuration file")
		rootCmd.PersistentFlags().StringSliceP("templates", "t", []string{}, "template(s) to apply")

		rootCmd.AddCommand(GetVersionCmd())
		rootCmd.AddCommand(GetListCmd())
		rootCmd.AddCommand(GetApplyCmd())
	}

	return rootCmd
}
