package newed

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/afero"

	"github.com/asphaltbuffet/newed/internal/config"
)

type List struct {
	cfg          *config.Config
	fs           afero.Fs
	templateDirs []string
}

// New creates a new List instance.
func New(cfg *config.Config, opts ...func(*List)) (*List, error) {
	l := &List{
		cfg:          cfg,
		fs:           afero.NewOsFs(),
		templateDirs: []string{},
	}

	for _, opt := range opts {
		opt(l)
	}

	return l, nil
}

// WithDirectory sets the directories to search for templates.
func WithDirectory(dirs ...string) func(*List) {
	return func(l *List) {
		l.templateDirs = append(l.templateDirs, dirs...)
	}
}

// GetTemplates lists the available templates.
func (l *List) GetTemplates(sub bool) (map[string][]string, error) {
	tmap := make(map[string][]string)

	for _, src := range l.templateDirs {
		if err := scanForTemplates(l.fs, src, tmap); err != nil {
			return nil, err
		}
	}

	if sub {
		if err := scanForSubTemplates(l.fs, tmap); err != nil {
			return nil, err
		}
	}

	return tmap, nil
}

func scanForTemplates(fs afero.Fs, src string, tmap map[string][]string) error {
	entries, err := afero.ReadDir(fs, src)
	if err != nil {
		return fmt.Errorf("reading directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			tmap[filepath.Join(src, entry.Name())] = []string{}
		}
	}

	return nil
}

func scanForSubTemplates(fs afero.Fs, tmap map[string][]string) error {
	for t := range tmap {
		subTmpls, err := afero.ReadDir(fs, t)
		if err != nil {
			return fmt.Errorf("reading subdirectory: %w", err)
		}

		for _, subTmpl := range subTmpls {
			if subTmpl.IsDir() {
				tmap[t] = append(tmap[t], subTmpl.Name())
			}
		}
	}

	return nil
}
