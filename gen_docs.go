//go:build ignore

package main

import (
	"os"

	"github.com/spf13/cobra/doc"

	"github.com/asphaltbuffet/newed/cmd"
)

func main() {
	rc := cmd.GetRootCommand()
	rc.InitDefaultCompletionCmd()
	rc.DisableAutoGenTag = true

	_ = os.Mkdir("./docs", 0755)

	_ = doc.GenMarkdownTree(rc, "./docs")
}
