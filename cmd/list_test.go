package cmd_test

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/asphaltbuffet/newed/cmd"
	"github.com/asphaltbuffet/newed/pkg/newed"
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

func TestPrintList(t *testing.T) {
	type args struct {
		tList map[string][]newed.Template
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty list",
			args: args{
				tList: map[string][]newed.Template{},
			},
			want: "",
		},
		{
			name: "one template",
			args: args{
				tList: map[string][]newed.Template{
					"/path/to/test": {
						{
							Name:     "test",
							Dir:      "/path/to/test",
							Base:     true,
							Sections: []string{"_base", "sub1"},
						},
					},
				},
			},
			want: "/path/to/test:\n    test [_base, sub1]\n",
		},
	}

	c := &cobra.Command{}
	b := bytes.NewBufferString("")
	c.SetOut(b)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd.PrintList(c, tt.args.tList)
			assert.Equal(t, tt.want, b.String())
		})
	}
}
