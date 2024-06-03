// Package main is the entry point for the CLI
package main

import (
	_ "github.com/spf13/cobra/doc"

	"github.com/asphaltbuffet/newed/cmd"
)

//go:generate go run gen_docs.go

func main() {
	cmd.Execute()
}
