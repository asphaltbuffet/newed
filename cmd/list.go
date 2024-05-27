package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/asphaltbuffet/newed/internal/config"
	"github.com/asphaltbuffet/newed/pkg/newed"
)

var (
	listCmd *cobra.Command
	showSub bool
)

func GetListCmd() *cobra.Command {
	if listCmd == nil {
		listCmd = &cobra.Command{
			Use:     "list",
			Aliases: []string{"l", "ls"},
			Short:   "list available templates",
			RunE:    runListCmd,
		}

		listCmd.Flags().BoolVarP(&showSub, "show-sub-templates", "s", false, "show sub-templates")
	}

	return listCmd
}

func runListCmd(cmd *cobra.Command, args []string) error {
	cf, _ := cmd.Flags().GetString("config-file")

	cfg, err := config.New(config.WithFile(cf))
	if err != nil {
		return err
	}

	tmpls := make(newed.Templates)
	dirs := []string{}

	for _, d := range args {
		var dir string
		dir, err = filepath.Abs(d)
		if err != nil {
			return err
		}

		var info fs.FileInfo
		info, err = os.Stat(dir)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			return fmt.Errorf("%s is not a directory", dir)
		}

		dirs = append(dirs, dir)
	}

	templateDirs := append(cfg.GetTemplateDirs(), dirs...)

	if err := tmpls.Load(templateDirs...); err != nil {
		return fmt.Errorf("loading templates: %w", err)
	}

	for k, v := range tmpls {
		var tmplOut string

		if showSub {
			tmplOut = fmt.Sprintf("%s [%s]", k, strings.Join(v.Sections, ", "))
		} else {
			tmplOut = k
		}

		cmd.Println(tmplOut)
	}

	return nil
}
