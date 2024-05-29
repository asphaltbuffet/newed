package newed_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func testTemplateSetup(t *testing.T) (func(), string, string) {
	t.Helper()
	testSrc := t.TempDir()
	testDest := t.TempDir()

	testFiles := []string{
		filepath.Join(testSrc, "fake_*.txt"),
		filepath.Join(testSrc, "fake-1", "_base", "fake_*.txt"),
		filepath.Join(testSrc, "fake-2", "_base", "fake_*.txt"),
		filepath.Join(testSrc, "fake-2", "fake2sub1", "fake_*.txt"),
		filepath.Join(testSrc, "fake-2", "fake2sub2", "fake_*.txt"),
	}

	for _, f := range testFiles {
		require.NoError(t, os.MkdirAll(filepath.Dir(f), 0o755))
		_, err := os.CreateTemp(filepath.Dir(f), filepath.Base(f))
		require.NoError(t, err)
	}

	require.NoError(t, os.MkdirAll(filepath.Join(testSrc, "fake-3"), 0o755))
	require.NoError(t, os.MkdirAll(filepath.Join(testSrc, "perm"), 0o200))

	return func() {
		// os.RemoveAll(testSrc) // this isn't necessary when using t.TempDir()
	}, testSrc, testDest
}
