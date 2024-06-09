package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var versionCmd *cobra.Command

func GetVersionCmd() *cobra.Command {
	if versionCmd == nil {
		versionCmd = &cobra.Command{
			Use:   "version",
			Short: "Print the version",
			Run: func(cmd *cobra.Command, _ []string) {
				sb := &strings.Builder{}

				if version != "" {
					sb.WriteString(version)
				}

				if revision != "" {
					sb.WriteString(fmt.Sprintf(" (%s)", revision))
				}

				cmd.Println(sb.String())
			},
		}
	}

	return versionCmd
}
