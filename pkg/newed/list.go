package newed

import (
	"fmt"

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
func (l *List) GetTemplates() error {
	for _, src := range l.templateDirs {
		// walk each template directory and print any directories
		entries, err := afero.ReadDir(l.fs, src)
		if err != nil {
			return fmt.Errorf("reading directory: %w", err)
		}

		for _, entry := range entries {
			if entry.IsDir() {
				fmt.Println(entry.Name())
			}
		}
	}

	return nil
}
