package cmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/asphaltbuffet/newed/cmd"
)

func TestGetListCmd(t *testing.T) {
	t.Run("create new instance", func(t *testing.T) {
		assert.NotNil(t, cmd.GetListCmd())
	})

	t.Run("return existing instance", func(t *testing.T) {
		lsCmd := cmd.GetListCmd()
		assert.Equal(t, lsCmd, cmd.GetListCmd())
	})
}
