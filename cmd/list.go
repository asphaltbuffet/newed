package cmd

import (
	"fmt"
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

	tmpls, err := newed.New(cfg)
	if err != nil {
		return err
	}

	if err = tmpls.Load(args...); err != nil {
		return err
	}

	// get list of directories
	dirs := make(map[string][]string)
	for name, template := range tmpls {
		d, _ := filepath.Abs(template.Dir)
		templateDir := filepath.Dir(d)

		if _, ok := dirs[templateDir]; !ok {
			dirs[templateDir] = []string{}
		}

		dirs[templateDir] = append(dirs[templateDir], name)
	}

	// print out the templates
	for dir, names := range dirs {
		cmd.Println(fmt.Sprintf("%s: ", dir))
		for _, name := range names {
			t := tmpls[name]

			sb := strings.Builder{}

			sb.WriteString("    ")
			sb.WriteString(t.Name)

			if t.Base {
				sb.WriteString("*")
			}

			sb.WriteString(fmt.Sprintf(" [%s]", strings.Join(t.Sections, ", ")))

			cmd.Println(sb.String())
		}
	}

	return nil
}
