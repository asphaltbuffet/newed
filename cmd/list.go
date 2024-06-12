package cmd

import (
	"fmt"
	"strings"

	"github.com/caarlos0/log"
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
			Use:     "list [-s|--show-sub-templates] [template dir]...",
			Aliases: []string{"l", "ls"},
			Short:   "list all templates",
			PreRun: func(cmd *cobra.Command, _ []string) {
				cf, err := cmd.Flags().GetString("config")
				if err != nil {
					log.WithError(err).Fatal("getting config filename")
				}

				log.WithField("file", cf).Info("loading configuration")

				cfg, err = config.New(config.WithFile(cf))
				if err != nil {
					log.WithError(err).Fatal("loading config file")
				}
			},
			RunE: runListCmd,
		}

		listCmd.Flags().BoolVarP(&showSub, "show-sub-templates", "s", false, "show sub-templates")
	}

	return listCmd
}

func runListCmd(cmd *cobra.Command, args []string) error {
	tColl, err := newed.New(cfg)
	if err != nil {
		return err
	}

	if err = tColl.Load(args...); err != nil {
		return err
	}

	PrintList(cmd, tColl.GetAllByDir())

	return nil
}

func PrintList(cmd *cobra.Command, tList map[string][]newed.Template) {
	// print out the templates
	for dir, templates := range tList {
		cmd.Println(fmt.Sprintf("%s:", dir))

		for _, t := range templates {
			sb := strings.Builder{}

			sb.WriteString("    ")
			sb.WriteString(t.Name)

			sb.WriteString(fmt.Sprintf(" [%s]", strings.Join(t.Sections, ", ")))

			cmd.Println(sb.String())
		}
	}
}
