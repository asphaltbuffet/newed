package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/asphaltbuffet/newed/internal/config"
	"github.com/asphaltbuffet/newed/pkg/newed"
)

var applyCmd *cobra.Command

func GetApplyCmd() *cobra.Command {
	if applyCmd == nil {
		applyCmd = &cobra.Command{
			Use:     "apply [flags] <path/to/dest>",
			Aliases: []string{},
			Args:    cobra.ExactArgs(1),
			Short:   "populate directory with template(s)",
			RunE:    runApplyCmd,
		}

		applyCmd.Flags().BoolP("dry-run", "n", false, "show what would be done without actually doing it")
		applyCmd.Flags().StringSliceP("templates", "t", []string{}, "template(s) to apply")
	}

	return applyCmd
}

func runApplyCmd(cmd *cobra.Command, args []string) error {
	cf, _ := cmd.Flags().GetString("config-file")

	cfg, err := config.New(config.WithFile(cf))
	if err != nil {
		return err
	}

	tmpls := make(newed.Templates)

	if err = tmpls.Load(cfg.GetTemplateDirs()...); err != nil {
		return fmt.Errorf("loading templates: %w", err)
	}

	templateFlags, err := cmd.Flags().GetStringSlice("templates")
	if err != nil {
		return err
	}

	isNoop, err := cmd.Flags().GetBool("dry-run")
	if err != nil {
		return err
	}

	if err = tmpls.Apply(templateFlags, args[0], isNoop); err != nil {
		return err
	}

	return nil
}
