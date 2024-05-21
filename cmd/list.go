package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/asphaltbuffet/newed/internal/config"
	"github.com/asphaltbuffet/newed/pkg/newed"
)

var listCmd *cobra.Command

func GetListCmd() *cobra.Command {
	if listCmd == nil {
		listCmd = &cobra.Command{
			Use:     "list",
			Aliases: []string{"l", "ls"},
			Args:    cobra.MinimumNArgs(1),
			Short:   "list available templates",
			RunE:    runListCmd,
		}
	}

	return listCmd
}

func runListCmd(cmd *cobra.Command, args []string) error {
	cf, _ := cmd.Flags().GetString("config-file")

	cfg, err := config.New(config.WithFile(cf))
	if err != nil {
		return err
	}

	dirs := []string{}

	for _, d := range args {
		var dir string
		dir, err = filepath.Abs(d)
		if err != nil {
			return fmt.Errorf("target directory: %w", err)
		}

		dirs = append(dirs, dir)
	}

	tmplList, err := newed.New(cfg, newed.WithDirectory(dirs...))
	if err != nil {
		return fmt.Errorf("creating list: %w", err)
	}

	return tmplList.GetTemplates()
}
