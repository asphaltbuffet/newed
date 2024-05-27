package newed

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/otiai10/copy"
)

// Apply copies the templates to the provided path.
func (t Templates) Apply(templates []string, path string, isNoop bool) error {
	dest, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	opts := getOpts(isNoop)

	expandedTemplates := Expand(templates...)

	for _, name := range expandedTemplates {
		template, ok := t[filepath.Dir(name)]
		if !ok {
			return fmt.Errorf("template %s not found", name)
		}

		// create full path to template directory
		tpath := filepath.Join(template.Dir, name)

		if err = copy.Copy(tpath, dest, *opts); err != nil {
			return err
		}
	}

	return nil
}

func getOpts(skipAll bool) *copy.Options {
	// by default we copy contents of symlink and merge directories and overwrite existing files
	opts := copy.Options{
		OnSymlink:   func(_ string) copy.SymlinkAction { return copy.Deep },
		OnDirExists: func(_ string, _ string) copy.DirExistsAction { return copy.Merge },
	}

	if skipAll {
		opts.Skip = func(_ fs.FileInfo, src string, dest string) (bool, error) {
			// skipping like this means we don't see contents in directories
			fmt.Printf("%s -> %s\n", src, dest)
			return true, nil
		}
	}

	return &opts
}
