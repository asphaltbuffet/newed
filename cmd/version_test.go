package cmd

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetVersionCmd(t *testing.T) {
	type args struct {
		version  string
		revision string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "not set",
			args: args{},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bytes.NewBufferString("")

			version = tt.args.version
			revision = tt.args.revision
			tcmd := GetVersionCmd()

			tcmd.SetOut(b)

			out, err := io.ReadAll(b)
			require.NoError(t, err)

			assert.Equal(t, tt.want, string(out))
		})
	}
}
