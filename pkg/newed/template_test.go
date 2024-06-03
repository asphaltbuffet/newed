package newed_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/asphaltbuffet/newed/pkg/newed"
)

func Test_Load(t *testing.T) {
	// set up test directories
	testTeardown, testSrc, _ := testTemplateSetup(t)
	defer testTeardown()

	testTemplates := make(newed.Templates)

	t.Run("add new template to empty", func(t *testing.T) {
		require.NoError(t, testTemplates.Load(testSrc))
		assert.Len(t, testTemplates, 3)
		assert.Contains(t, testTemplates, "fake-1")
		assert.Contains(t, testTemplates, "fake-2")
		assert.Contains(t, testTemplates, "fake-3")
	})

	t.Run("load invalid dir", func(t *testing.T) {
		testTemplates = make(newed.Templates)
		require.NoError(t, testTemplates.Load("bad-dir"))
		assert.Empty(t, testTemplates)
	})
}

func Test_Add(t *testing.T) {
	// set up test directories
	testTeardown, testSrc, _ := testTemplateSetup(t)
	defer testTeardown()

	fakeFile, _ := os.CreateTemp(testSrc, "fakeFile_*.txt")

	testTemplates := make(newed.Templates)

	t.Run("add new template to empty", func(t *testing.T) {
		require.NoError(t, testTemplates.Add(filepath.Join(testSrc, "fake-1")))
		assert.Len(t, testTemplates, 1)
		require.Contains(t, testTemplates, "fake-1")
		assert.True(t, testTemplates["fake-1"].Base)
		assert.Len(t, testTemplates["fake-1"].Sections, 1)
	})

	t.Run("add existing template", func(t *testing.T) {
		require.NoError(t, testTemplates.Add(filepath.Join(testSrc, "fake-1")))
		assert.Len(t, testTemplates, 1)
	})

	t.Run("add new template to others present", func(t *testing.T) {
		require.NoError(t, testTemplates.Add(filepath.Join(testSrc, "fake-2")))
		assert.Len(t, testTemplates, 2)
		require.Contains(t, testTemplates, "fake-2")
		assert.True(t, testTemplates["fake-2"].Base)
		assert.NotEmpty(t, testTemplates["fake-2"].Sections)
	})

	t.Run("add template without base", func(t *testing.T) {
		require.NoError(t, testTemplates.Add(filepath.Join(testSrc, "fake-3")))
		assert.Len(t, testTemplates, 3)
		require.Contains(t, testTemplates, "fake-3")
		assert.False(t, testTemplates["fake-3"].Base)
		assert.Empty(t, testTemplates["fake-3"].Sections)
	})

	t.Run("fail to add non-existant template", func(t *testing.T) {
		require.Error(t, testTemplates.Add(filepath.Join(testSrc, "no-fake-here")))
		assert.Len(t, testTemplates, 3)
		require.NotContains(t, testTemplates, "no-fake-here")
	})

	t.Run("fail to add non-directory", func(t *testing.T) {
		err := testTemplates.Add(fakeFile.Name())
		require.ErrorIs(t, err, os.ErrInvalid)

		assert.Len(t, testTemplates, 3)
	})

	t.Run("fail due to permissions", func(t *testing.T) {
		err := testTemplates.Add(filepath.Join(testSrc, "perm"))
		require.Error(t, err)
		assert.ErrorIs(t, err, os.ErrPermission)
	})
}

func Test_Expand(t *testing.T) {
	type args struct {
		templates []string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "single base template",
			args: args{
				templates: []string{"fake"},
			},
			want: []string{"fake/_base"},
		},
		{
			name: "single with additions",
			args: args{
				templates: []string{"fake+fake_add1+fake_add2"},
			},
			want: []string{"fake/_base", "fake/fake_add1", "fake/fake_add2"},
		},
		{
			name: "two base templates",
			args: args{
				templates: []string{"fake", "fake2"},
			},
			want: []string{"fake/_base", "fake2/_base"},
		},
		{
			name: "two base templates with additions",
			args: args{
				templates: []string{"fake+fake_add1+fake_add2", "fake2+fake2_add1"},
			},
			want: []string{"fake/_base", "fake/fake_add1", "fake/fake_add2", "fake2/_base", "fake2/fake2_add1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, newed.Expand(tt.args.templates...))
		})
	}
}

func TestTemplates_GetAllByDir(t *testing.T) {
	tests := []struct {
		name string
		tr   newed.Templates
		want map[string][]newed.Template
	}{
		{
			name: "no templates",
			tr:   map[string]newed.Template{},
			want: map[string][]newed.Template{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.tr.GetAllByDir())
		})
	}
}
