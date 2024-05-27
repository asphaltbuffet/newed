package cmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/asphaltbuffet/newed/cmd"
)

func Test_GetApplyCmd(t *testing.T) {
	t.Run("create new instance", func(t *testing.T) {
		assert.NotNil(t, cmd.GetApplyCmd())
	})

	t.Run("return existing instance", func(t *testing.T) {
		apCmd := cmd.GetApplyCmd()
		assert.Equal(t, apCmd, cmd.GetApplyCmd())
	})
}
