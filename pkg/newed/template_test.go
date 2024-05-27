package newed

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTemplates_Add(t *testing.T) {
	// set up test directories
	testSrc := t.TempDir()
	testDest := t.TempDir()
	require.NotEqual(t, testSrc, testDest)

	os.MkdirAll(filepath.Join(testSrc, "fake-1", "_base"), 0o755)
	os.CreateTemp(filepath.Join(testSrc, "fake-1", "_base"), "fake_*.txt")

	os.MkdirAll(filepath.Join(testSrc, "fake-2", "_base"), 0o755)
	os.CreateTemp(filepath.Join(testSrc, "fake-2", "_base"), "fake_*.txt")

	os.MkdirAll(filepath.Join(testSrc, "fake-2", "fake2sub1"), 0o755)
	os.CreateTemp(filepath.Join(testSrc, "fake-2", "fake2sub1"), "fake_*.txt")

	os.MkdirAll(filepath.Join(testSrc, "fake-2", "fake2sub2"), 0o755)
	os.CreateTemp(filepath.Join(testSrc, "fake-2", "fake2sub2"), "fake_*.txt")

	os.MkdirAll(filepath.Join(testSrc, "fake-3"), 0o755)

	test_templates := make(Templates)

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
}
