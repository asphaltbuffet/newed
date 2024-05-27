// Package cmd contains all CLI commands used by the application.
package cmd_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/asphaltbuffet/newed/cmd"
)

func Test_GetRootCommand(t *testing.T) {
	t.Run("create new instance of root command", func(t *testing.T) {
		assert.NotNil(t, cmd.GetRootCommand())
	})

	t.Run("return existing instance of root command", func(t *testing.T) {
		rootCmd := cmd.GetRootCommand()
		assert.Equal(t, rootCmd, cmd.GetRootCommand())

		rootCmd.Version = "fake"
		assert.Equal(t, "fake", cmd.GetRootCommand().Version)
	})
}

func Test_Execute(t *testing.T) {
	rootcmd := cmd.GetRootCommand()
	b := bytes.NewBufferString("")
	rootcmd.SetOut(b)
	rootcmd.SetErr(b)

	cmd.Execute()

	out, err := io.ReadAll(b)
	require.NoError(t, err)

	assert.NotEmpty(t, out)
}
