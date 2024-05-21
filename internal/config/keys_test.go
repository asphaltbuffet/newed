package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigKey_String(t *testing.T) {
	tests := []struct {
		name string
		key  Key
		want string
	}{
		{"one word", LogLevelKey, "logging"},
		{"nested", TemplateDirsKey, "template.dirs"},
		{"unknown", "fake", "fake"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.key.String())
		})
	}
}
