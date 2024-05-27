package newed_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/asphaltbuffet/newed/pkg/newed"
)

func TestTemplates_Apply(t *testing.T) {
	// set up test directories
	testTeardown, testSrc, testDest := testTemplateSetup(t)
	defer testTeardown()

	type args struct {
		tCollection newed.Templates
		selected    []string
		dest        string
		noop        bool
	}

	tests := []struct {
		name      string
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "apply only template in collection",
			args: args{
				tCollection: map[string]newed.Template{
					"fake-1": {
						Name:     "fake-1",
						Dir:      testSrc,
						Base:     true,
						Sections: []string{"_base"},
					},
				},
				selected: []string{"fake-1"},
				dest:     testDest,
			},
			assertion: assert.NoError,
		},
		{
			name: "apply invalid template",
			args: args{
				tCollection: map[string]newed.Template{
					"fake-1": {
						Name:     "fake-1",
						Dir:      testSrc,
						Base:     true,
						Sections: []string{"_base"},
					},
				},
				selected: []string{"fake-invalid"},
				dest:     testDest,
			},
			assertion: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tcoll := tt.args.tCollection
			tt.assertion(t, tcoll.Apply(tt.args.selected, tt.args.dest, tt.args.noop))
		})
	}
}
