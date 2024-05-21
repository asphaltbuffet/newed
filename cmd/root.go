// Package cmd contains all CLI commands used by the application.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/asphaltbuffet/newed/internal/config"
)

// application build information set by the linker.
var (
	Version string
	Date    string
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
	var cfgFile string
	if rootCmd == nil {
		rootCmd = &cobra.Command{
			Use:     "newed",
			Version: fmt.Sprintf("%s\n%s", Version, Date),
			Short:   "newed creates projects from templates",
			Run: func(cmd *cobra.Command, _ []string) {
				cfg, err := config.New(config.WithFile(cfgFile))
				if err != nil {
					cmd.PrintErr(err)
				}

				cmd.Println("config file:", cfg.GetConfigFileUsed())
			},
		}
	}

	rootCmd.Flags().StringVarP(&cfgFile, "config-file", "c", "", "configuration file")

	rootCmd.AddCommand(GetListCmd())

	return rootCmd
}
